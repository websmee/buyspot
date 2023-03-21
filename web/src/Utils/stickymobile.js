export default {
    addPlatformClass: () => {
        document.addEventListener('DOMContentLoaded', () => {
            const iOSVersion = () => {
                let d, v;
                if (/iP(hone|od|ad)/.test(navigator.platform)) {
                    v = (navigator.appVersion).match(/OS (\d+)_(\d+)_?(\d+)?/);
                    d = { status: true, version: parseInt(v[1], 10), info: parseInt(v[1], 10) + '.' + parseInt(v[2], 10) + '.' + parseInt(v[3] || 0, 10) };
                } else { d = { status: false, version: false, info: '' } }
                return d;
            }
            if (iOSVersion().version > 14) { document.querySelectorAll('#page')[0].classList.add('min-ios15'); }
        });
    },

    getMenuOpenListener: (dataMenuId) => {
        return () => {
            //Close Existing Opened Menus
            const activeMenu = document.querySelectorAll('.menu-active');
            for (let i = 0; i < activeMenu.length; i++) { activeMenu[i].classList.remove('menu-active'); }
            //Open Clicked Menu
            document.getElementById(dataMenuId).classList.add('menu-active');
            document.getElementsByClassName('menu-hider')[0].classList.add('menu-active');
            //Check and Apply Effects
            var menu = document.getElementById(dataMenuId);
            var menuEffect = menu.getAttribute('data-menu-effect');
            var menuLeft = menu.classList.contains('menu-box-left');
            var menuRight = menu.classList.contains('menu-box-right');
            var menuTop = menu.classList.contains('menu-box-top');
            var menuBottom = menu.classList.contains('menu-box-bottom');
            var menuWidth = menu.offsetWidth;
            var menuHeight = menu.offsetHeight;

            if (menuEffect === "menu-push") {
                var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
                var menuWidth = document.getElementById(dataMenuId).getAttribute('data-menu-width');
                if (menuLeft) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(" + menuWidth + "px)" } }
                if (menuRight) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + menuWidth + "px)" } }
                if (menuBottom) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(-" + menuHeight + "px)" } }
                if (menuTop) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(" + menuHeight + "px)" } }
            }
            if (menuEffect === "menu-parallax") {
                var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
                var menuWidth = document.getElementById(dataMenuId).getAttribute('data-menu-width');
                if (menuLeft) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(" + menuWidth / 10 + "px)" } }
                if (menuRight) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + menuWidth / 10 + "px)" } }
                if (menuBottom) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(-" + menuHeight / 5 + "px)" } }
                if (menuTop) { for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateY(" + menuHeight / 5 + "px)" } }
            }
        }
    },

    getMenuCloseListener: () => {
        return () => {
            const activeMenu = document.querySelectorAll('.menu-active');
            for (let i = 0; i < activeMenu.length; i++) { activeMenu[i].classList.remove('menu-active'); }
            var wrappers = document.querySelectorAll('.header, #footer-bar, .page-content');
            for (let i = 0; i < wrappers.length; i++) { wrappers[i].style.transform = "translateX(-" + 0 + "px)" }
            var iframes = document.querySelectorAll('iframe');
            iframes.forEach(el => { var hrefer = el.getAttribute('src'); el.setAttribute('newSrc', hrefer); el.setAttribute('src', ''); var newSrc = el.getAttribute('newSrc'); el.setAttribute('src', newSrc) });
        }
    },

    bindMenu: (dataMenuId, menuOpenListener, menuCloseListener) => {
        document.getElementById(dataMenuId).style.display = 'block';

        document.querySelectorAll('[data-menu=' + dataMenuId + ']').forEach(el => {
            el.addEventListener('click', menuOpenListener);
        });

        document.querySelectorAll('.close-menu, .menu-hider').forEach(el => el.addEventListener('click', menuCloseListener));
    },

    unbindMenu: (dataMenuId, menuOpenListener, menuCloseListener) => {
        document.querySelectorAll('[data-menu=' + dataMenuId + ']').forEach(el => {
            el.removeEventListener('click', menuOpenListener);
        });

        document.querySelectorAll('.close-menu, .menu-hider').forEach(el => el.removeEventListener('click', menuCloseListener));
    },

    showPreloader: () => {
        document.getElementById('preloader').classList.remove('preloader-hide')
    },

    hidePreloader: () => {
        document.getElementById('preloader').classList.add('preloader-hide')
    },
};

const preventDefault = (event) => {
    event.preventDefault();
    return false;
};