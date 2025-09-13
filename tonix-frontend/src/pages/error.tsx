type ErrorPageProps = {
    message: string,
    code: number,
}

export const ErrorPage = ({message, code}: ErrorPageProps) => {
    return <div className="font-main text-warn-hl">
        {code}: {message}
    </div>
}
