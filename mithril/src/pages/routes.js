import homePage from "./home-page";
import personPage from "./person-page";

var routes = {
  home: "/home",
  person: "/person",
  routeMap: {},
};

routes.routeMap[routes.home] = homePage;
routes.routeMap[routes.person] = personPage;
export default routes;
