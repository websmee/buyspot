const iOSversion = () => {
    let d, v;
    if (/iP(hone|od|ad)/.test(navigator.platform)) {
        v = (navigator.appVersion).match(/OS (\d+)_(\d+)_?(\d+)?/);
        d = { status: true, version: parseInt(v[1], 10), info: parseInt(v[1], 10) + '.' + parseInt(v[2], 10) + '.' + parseInt(v[3] || 0, 10) };
    } else { d = { status: false, version: false, info: '' } }
    return d;
}

const openListener = e => {
    //Close Existing Opened Menus
    const activeMenu = document.querySelectorAll('.menu-active');
    for (let i = 0; i < activeMenu.length; i++) { activeMenu[i].classList.remove('menu-active'); }
    //Open Clicked Menu
    var menuData = e.currentTarget.getAttribute('data-menu');
    document.getElementById(menuData).classList.add('menu-active');
    document.getElementsByClassName('menu-hider')[0].classList.add('menu-active');
    //Check and Apply Effects
    var menu = document.getElementById(menuData);
    var menuEffect = menu.getAttribute('data-menu-effect');
    var menuLeft = menu.classList.contains('menu-box-left');
    var menuRight = menu.classList.contains('menu-box-right');
    var menuTop = menu.classList.contains('menu-box-top');
    var menuBottom = menu.classList.contains('menu-box-bottom');
    var menuWidth = menu.offsetWidth;
    var menuHeight = menu.offsetHeight;

    if (menuEffect === "menu-push") {
        var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
        var menuWidth = document.getElementById(menuData).getAttribute('data-menu-width');
        if (menuLeft) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(" + menuWidth + "px)" } }
        if (menuRight) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + menuWidth + "px)" } }
        if (menuBottom) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(-" + menuHeight + "px)" } }
        if (menuTop) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(" + menuHeight + "px)" } }
    }
    if (menuEffect === "menu-parallax") {
        var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
        var menuWidth = document.getElementById(menuData).getAttribute('data-menu-width');
        if (menuLeft) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(" + menuWidth / 10 + "px)" } }
        if (menuRight) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + menuWidth / 10 + "px)" } }
        if (menuBottom) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(-" + menuHeight / 5 + "px)" } }
        if (menuTop) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(" + menuHeight / 5 + "px)" } }
    }
}

const closeListener = e => {
    const activeMenu = document.querySelectorAll('.menu-active');
    for (let i = 0; i < activeMenu.length; i++) { activeMenu[i].classList.remove('menu-active'); }
    var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
    for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + 0 + "px)" }
    var iframes = document.querySelectorAll('iframe');
    iframes.forEach(el => { var hrefer = el.getAttribute('src'); el.setAttribute('newSrc', hrefer); el.setAttribute('src', ''); var newSrc = el.getAttribute('newSrc'); el.setAttribute('src', newSrc) });
}

const preventDefault = (event) => {
    event.preventDefault();
    return false;
};

const bindMenus = () => {
    document.querySelectorAll('.menu').forEach(el => { el.style.display = 'block' });
    var menus = document.querySelectorAll('.menu');
    if (menus.length) {
        var menuSidebar = document.querySelectorAll('.menu-box-left, .menu-box-right');
        menuSidebar.forEach(function (e) {
            if (e.getAttribute('data-menu-width') === "cover") {
                e.style.width = '100%'
            } else {
                e.style.width = (e.getAttribute('data-menu-width')) + 'px'
            }
        })
        var menuSheets = document.querySelectorAll('.menu-box-bottom, .menu-box-top, .menu-box-modal');
        menuSheets.forEach(function (e) {
            if (e.getAttribute('data-menu-width') === "cover") {
                e.style.width = '100%'
                e.style.height = '100%'
            } else {
                e.style.width = (e.getAttribute('data-menu-width')) + 'px'
                e.style.height = (e.getAttribute('data-menu-height')) + 'px'
            }
        })

        //Opening Menus
        var menuOpen = document.querySelectorAll('[data-menu]');
        menuOpen.forEach(el => el.addEventListener('click', openListener));

        //Closing Menus
        const menuClose = document.querySelectorAll('.close-menu, .menu-hider');
        menuClose.forEach(el => el.addEventListener('click', closeListener));
    }
}

const bindEmptyLinks = () => {
    document.querySelectorAll('a[href="#"]').forEach(el => el.addEventListener('click', preventDefault));
}

const unbindMenus = () => {
    var menus = document.querySelectorAll('.menu');
    if (menus.length) {
        //Opening Menus
        var menuOpen = document.querySelectorAll('[data-menu]');
        menuOpen.forEach(el => el.removeEventListener('click', openListener));

        //Closing Menus
        const menuClose = document.querySelectorAll('.close-menu, .menu-hider');
        menuClose.forEach(el => el.removeEventListener('click', closeListener));
    }
}

const unbindEmptyLinks = () => {
    document.querySelectorAll('a[href="#"]').forEach(el => el.removeEventListener('click', preventDefault));
}

const bind = () => {
    bindMenus();
    bindEmptyLinks();
}

const unbind = () => {
    unbindMenus();
    unbindEmptyLinks();
}

export { bind, unbind, iOSversion }