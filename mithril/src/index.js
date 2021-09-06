import m from "mithril";
import homePage from "./pages/home-page";
import personPage from "./pages/person-page";
import urls from "./pages/urls";

var root = document.body;
var routes = {
  routeMap: {},
};
routes.routeMap[urls.home] = homePage;
routes.routeMap[urls.person] = personPage;

m.route(root, urls.home, routes.routeMap);
