import './globals.css'
// import type { Metadata } from 'next'
// import { Inter } from 'next/font/google'
import { Navigation } from './components/navigation/navigation'

import { PropsWithChildren } from "react";
import Link from 'next/link';

export const Layout = ({ children }: PropsWithChildren) => {
    return (
        <div>
            <div className="drawer drawer-mobile">
                <LeftSidebar />

            </div>

            <Navigation />
            <main className='container mx-auto'>
                {children}
            </main>
        </div>
    );
}


function LeftSidebar() {
    // const location = useLocation();

    // const dispatch = useDispatch()


    const close = (e) => {
        document.getElementById('left-sidebar-drawer').click()
    }

    return (
        <div className="drawer-side ">
            <label htmlFor="left-sidebar-drawer" className="drawer-overlay"></label>
            <ul className="menu  pt-2 w-80 bg-base-100 text-base-content">
                <button className="btn btn-ghost bg-base-300  btn-circle z-50 top-0 right-0 mt-4 mr-2 absolute lg:hidden" onClick={() => close()}>
                </button>

                <li className="mb-2 font-semibold text-xl">

                    <Link href={'/app/welcome'}><img className="mask mask-squircle w-10" src="/logo192.png" alt="DashWind Logo" />DashWind</Link> </li>
                <li className="" key={1}>



                </li>

            </ul>
        </div>
    )
}


export default Layout;