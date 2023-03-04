import Footer from "Layouts/footer"
import { Outlet } from "react-router-dom"

function Layout() {
    return (
        <>
            <Outlet />
            <Footer />
        </>
    )
}

export default Layout