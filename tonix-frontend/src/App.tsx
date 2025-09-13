import { BrowserRouter, Route, Routes } from 'react-router'
import './App.css'
import { ErrorPage } from './pages/error'
import { RegistrationPage } from './pages/registration'

function App() {
    return (
        <>
            <BrowserRouter>
                <Routes>
                    <Route path='/registration' element={<RegistrationPage />} />
                    <Route path='*' element={<ErrorPage message='Страница не найдена' code={404} />} />
                </Routes>
            </BrowserRouter>
        </>
    )
}

export default App
