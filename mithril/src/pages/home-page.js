import m from "mithril";
import urls from "./urls";
import deafaultLayout from "./default-layout";
import { CreateQueryString } from "./util";
import defaultLayout from "./default-layout";

const component = {
  view: function () {
    return m("div", [
      m("h1", "Home page"),
      m(
        m.route.Link,
        { href: urls.person + CreateQueryString({ id: "colcord.joanna" }) },
        urls.person
      ),
    ]);
  },
};

export default defaultLayout(component);
