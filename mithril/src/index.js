import m from "mithril";
import homePage from "./pages/home-page";
import routes from "./pages/routes";

var root = document.body;
m.route(root, routes.home, routes.routeMap);
