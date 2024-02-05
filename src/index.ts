// import * as controls from "./button";
import 'htmx.org'

// @ts-ignore
window.htmx = require('htmx.org')

// @ts-ignore
interface HtmxAfterRequestEvent extends CustomEvent {
    detail: {
        pathInfo: {
            requestPath: string
        };
    };
}

// update history state + browser url after htx-get -- for page nav
// @ts-ignore
document.body.addEventListener('htmx:afterRequest', function(e: HtmxAfterRequestEvent) {
    const url = e.detail.pathInfo.requestPath
    console.info("htmx afterRequest: pushing history state: " + url)
    history.pushState(null, '', url);
});