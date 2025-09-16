package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	"tonix/backend/api/dto/view"
	"tonix/backend/api/utils"
	"tonix/backend/model"

	"github.com/saryginrodion/stackable"
)

const ALLOWED_MIMETYPES string = "audio/ogg;audio/wav;audio/mpeg;image/jpeg;image/png;/image/gif"

var UploadFile = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		// Reading file
		if err := ctx.Request.ParseMultipartForm(int64(ctx.Shared.Environment.UPLOADS_MAX_SIZE_MB) << 20); err != nil {
			return err
		}

		file, header, err := ctx.Request.FormFile("file")

		if err != nil {
			return err
		}
		defer file.Close()

		// Validating mimetype
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			return dto.NewApiError(http.StatusBadRequest, "Can not detect mimetype of file due to it's small size (<512b).")
		}

		mimetype := http.DetectContentType(buffer)
		if !strings.Contains(ALLOWED_MIMETYPES, mimetype) {
			return dto.NewApiError(http.StatusBadRequest, "This mimetype is not allowed. Allowed mimetypes: "+ALLOWED_MIMETYPES)
		}

		file.Seek(0, io.SeekStart)

		// Writing file
		uploadTime := time.Now()
		dirName := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			uploadTime.Year(), uploadTime.Month(), uploadTime.Day(),
			uploadTime.Hour(), uploadTime.Minute(), uploadTime.Second())
		absoluteDirPath := filepath.Join(ctx.Shared.Environment.UPLOADS_DIRECTORY, dirName)

		fileName := header.Filename

		os.Mkdir(absoluteDirPath, 0744)
		dst, err := os.Create(filepath.Join(absoluteDirPath, fileName))
		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err = io.Copy(dst, file); err != nil {
			return err
		}


		// Adding to file DB
		files := model.Files(ctx.Shared.DB)
		uploadedFileId, err := files.AddFile(&model.File{
			AuthorId: ctx.Local.AccessJWT.Payload.Data.Uid,
			Path:     filepath.Join(dirName, fileName),
			Mimetype: mimetype,
			Filename: fileName,
		})

		if err != nil {
			return err
		}

		// Returning data of new file
		ctx.Response, _ = stackable.JsonResponse(http.StatusOK, view.FileId{Id: *uploadedFileId})

		return next()
	},
)

var ReadFile = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		fileId := ctx.Request.PathValue("id")
		files := model.Files(ctx.Shared.DB)

		fileData, err := files.ById(fileId)
		if err != nil {
			return err
		}

		filePath := filepath.Join(ctx.Shared.Environment.UPLOADS_DIRECTORY, fileData.Path)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("here 1")
			return err
		}
		fileStat, err := file.Stat()

		headers := stackable.NewHeadersContainer()
		headers.Set("Content-Type", fileData.Mimetype)
		headers.Set("Content-Length", strconv.Itoa(int(fileStat.Size())))

		ctx.Response = stackable.NewHttpResponseRaw(
			headers,
			http.StatusOK,
			utils.CreateAutoclosingFile(file),
		)

		return next()
	},
)
