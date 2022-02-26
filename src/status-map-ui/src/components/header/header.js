import React from 'react';
import './header.css';

export default () => {
    return (
        <header>
            <nav className="navbar navbar-dark navbar-expand-md py-3 px-4">
                <a className="navbar-brand mr-4" href="https://odj.cloud">
                    <svg className="mr-2" width="24" height="26" viewBox="0 0 24 26" fill="none"
                        xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                        <rect y="0.5" width="24" height="25" fill="url(#pattern0)" />
                        <defs>
                            <pattern id="pattern0" patternContentUnits="objectBoundingBox" width="1" height="1">
                                <use xlinkHref="#image0" transform="scale(0.0416667 0.04)" />
                            </pattern>
                            <image id="image0" width="24" height="25"
                                xlinkHref="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAZCAYAAAArK+5dAAAAAXNSR0IArs4c6QAAADhlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAAqACAAQAAAABAAAAGKADAAQAAAABAAAAGQAAAAD8nX4nAAACC0lEQVRIDaWUPU7DQBBGEwRISEhYSKCULmkQ3AAXFHTACTCcgCP4BpgTODUNlHTODYCWBjgBQaIP7zOzluOsvUGM9DK78/ON19lkMAjYbDaLoYAnkH3CPVwEWsNpRFKQTeEBMsjhHWQlRGElTwWNTnzsE7G8ButkfxuiBtCrGHtm1yHyZ9C0gk3/MAqOoARZXKt1LKjREJ02B5n/RCRUqKd2Nu7Q7AzTeGjN870EM0tM8An0H7NzxGBAr07yWZeYIK7/fdcNgQU61cO6shUW1/Bh3sX/7Rl0VomwCN4W3zT6TqGAEqSxpzp8BM8W21RAlvlEfDFqE3iDtiWunkRsyRO9oi+IXbLP03RDvoS4r244HL5bfrTKYgI6bkRiaokFR74gmFri1rz2W7aunbRs861X5O5u2UgovgYFtC13SiQyl3QxeWKFxUdVnE0KUwtW3z7rR9u3XXNAbskHE75g/2axVLGhPmQEY1wKz7AGd9BlbkhKQQSXVqjX+AIZr7saavF513gKe5hepweqjKptt+70FCW9covJpFOMhK5p21IL6Pp+tJOefX0CT24xxAOuQwR7oC8/ZMmiSiCC4g7o568rPIE+2wjIzadRWoVX0LV9ghh0kjG07Xi+e4kdCpsg0aZpUNIIaPj5EnL+EppHIJEcZLtwUK1+/yn3/Z1/jCJ4CIlrY30F+mEtbT/SJSOBmKzEKAAAAABJRU5ErkJggg==" />
                        </defs>
                    </svg>
                    ODJ
                </a>
                <div className="collapse navbar-collapse" id="navbarNav">
                    <ul className="navbar-nav w-100">
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/dashboard">Dashboard</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/myworkspace">My workspace</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/subscriptions">Subscriptions</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/prods/all">Products</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/service_catalog/?org=sit&amp;runtime=k8s">Service Catalog</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link active" href="/">Status Map</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="https://odj.cloud/docs/">Docs</a>
                        </li>
                    </ul>
                </div>
            </nav>
        </header>
    );
}
