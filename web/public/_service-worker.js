// To clear cache on devices, always increase APP_VER number after making changes.
// The app will serve fresh content right away or after 2-3 refreshes (open / close)
var APP_NAME = 'StickyMobile';
var APP_VER = '4.8.1';
var CACHE_NAME = APP_NAME + '-' + APP_VER;

// Files required to make this app work offline.
// Add all files you want to view offline below.
// Leave REQUIRED_FILES = [] to disable offline.
var REQUIRED_FILES = [
	// HTML Files
	'index.html',
	// Styles
	'stickymobile/styles/style.css',
	'stickymobile/styles/bootstrap.css',
	// Scripts
	'stickymobile/scripts/main.js',
	'stickymobile/scripts/bootstrap.min.js',
	// Plugins
	// 'stickymobile/plugins/before-after/before-after.css',
	// 'stickymobile/plugins/before-after/before-after.js',
	'stickymobile/plugins/charts/charts.js',
	// 'stickymobile/plugins/charts/charts-call-graphs.js',
	// 'stickymobile/plugins/countdown/countdown.js',
	// 'stickymobile/plugins/filterizr/filterizr.js',
	// 'stickymobile/plugins/filterizr/filterizr.css',
	// 'stickymobile/plugins/filterizr/filterizr-call.js',
	// 'stickymobile/plugins/galleryViews/gallery-views.js',
	// 'stickymobile/plugins/glightbox/glightbox.js',
	// 'stickymobile/plugins/glightbox/glightbox.css',
	// 'stickymobile/plugins/glightbox/glightbox-call.js',
	// Fonts
	'stickymobile/fonts/css/fontawesome-all.min.css',
	'stickymobile/fonts/webfonts/fa-brands-400.woff2',
	'stickymobile/fonts/webfonts/fa-regular-400.woff2',
	'stickymobile/fonts/webfonts/fa-solid-900.woff2',
	// Images
	'stickymobile/images/empty.png',
];

// Service Worker Diagnostic. Set true to get console logs.
var APP_DIAG = false;

//Service Worker Function Below.
self.addEventListener('install', function(event) {
	event.waitUntil(
		caches.open(CACHE_NAME)
		.then(function(cache) {
			//Adding files to cache
			return cache.addAll(REQUIRED_FILES);
		}).catch(function(error) {
			//Output error if file locations are incorrect
			if(APP_DIAG){console.log('Service Worker Cache: Error Check REQUIRED_FILES array in _service-worker.js - files are missing or path to files is incorrectly written -  ' + error);}
		})
		.then(function() {
			//Install SW if everything is ok
			return self.skipWaiting();
		})
		.then(function(){
			if(APP_DIAG){console.log('Service Worker: Cache is OK');}
		})
	);
	if(APP_DIAG){console.log('Service Worker: Installed');}
});

self.addEventListener('fetch', function(event) {
	event.respondWith(
		//Fetch Data from cache if offline
		caches.match(event.request)
			.then(function(response) {
				if (response) {return response;}
				return fetch(event.request);
			}
		)
	);
	if(APP_DIAG){console.log('Service Worker: Fetching '+APP_NAME+'-'+APP_VER+' files from Cache');}
});

self.addEventListener('activate', function(event) {
	event.waitUntil(self.clients.claim());
	event.waitUntil(
		//Check cache number, clear all assets and re-add if cache number changed
		caches.keys().then(cacheNames => {
			return Promise.all(
				cacheNames
					.filter(cacheName => (cacheName.startsWith(APP_NAME + "-")))
					.filter(cacheName => (cacheName !== CACHE_NAME))
					.map(cacheName => caches.delete(cacheName))
			);
		})
	);
	if(APP_DIAG){console.log('Service Worker: Activated')}
});