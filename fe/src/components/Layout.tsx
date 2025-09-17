import React, {type ReactNode} from "react";
import {Link, useLocation} from "react-router";
import {routes} from "../utils/routes";

interface SidebarLayoutProps {
  children: ReactNode;
}

const Layout: React.FC<SidebarLayoutProps> = ({children}) => {
  const location = useLocation();
  const menuItems = [
    {name: "Dashboard", path: routes.index},
    {name: "Profile", path: "/profile"},
    {name: "Settings", path: "/settings"},
  ];
  return (
    <div className="flex h-screen bg-base-200">
      <div className="w-64 bg-base-100 border-r border-base-300 flex flex-col">
        <div className="p-6 text-2xl font-bold border-b border-base-300">MyApp</div>

        <ul className="menu p-4 flex-1 gap-2">
          {menuItems.map((item) => (
            <li key={item.name} className={location.pathname === item.path ? "bg-primary text-primary-content rounded-lg" : ""}>
              <Link to={item.path}>{item.name}</Link>
            </li>
          ))}
        </ul>

        <div className="p-6 border-t border-base-300 relative">
          <div className="dropdown dropdown-top">
            <div tabIndex={0} role="button" className="btn m-1">
              <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
                <div className="w-12 rounded-full">
                  <img src="https://i.pravatar.cc/300" alt="User Avatar" />
                </div>
              </label>
            </div>
            <ul tabIndex={0} className="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm">
              <li>
                <a>Item 1</a>
              </li>
              <li>
                <a>Item 2</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <div className="flex-1 p-6 overflow-auto">{children}</div>
    </div>
  );
};

export default Layout;
