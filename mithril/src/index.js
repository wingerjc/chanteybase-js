import m from "mithril";
import homePage from "./pages/home-page";
import routes from "./pages/routes";

var root = document.body;
var routing = {};
routing[routes.home] = homePage;

console.log(routing);
m.route(root, routes.home, routing);
