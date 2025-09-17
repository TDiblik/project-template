import React, {type ReactNode} from "react";
import {Link, useLocation, useNavigate} from "react-router";
import {routes} from "../utils/routes";

interface SidebarLayoutProps {
  children: ReactNode;
}

const Layout: React.FC<SidebarLayoutProps> = ({children}) => {
  const location = useLocation();
  const navigate = useNavigate();

  const menuItems = [
    {name: "Dashboard", path: routes.index},
    {name: "Profile", path: routes.profile},
    {name: "Settings", path: routes.settings},
  ];

  return (
    <div className="flex h-screen bg-base-200">
      <div className="w-64 bg-base-100 border-r border-base-300 flex flex-col">
        <div className="p-6 text-2xl font-bold border-b border-base-300">project-template</div>

        {/* Menu */}
        <ul className="menu p-4 flex-1 gap-2">
          {menuItems.map((item) => (
            <li key={item.name} className={location.pathname.startsWith(item.path) ? "bg-primary text-primary-content rounded-lg" : ""}>
              <Link to={item.path}>{item.name}</Link>
            </li>
          ))}
        </ul>

        {/* Avatar Dropdown */}
        <div className="p-6 border-t border-base-300">
          <div className="dropdown dropdown-top">
            <div tabIndex={0} className="cursor-pointer">
              <div className="btn btn-ghost btn-circle avatar">
                <div className="w-12 rounded-full">
                  <img src="https://i.pravatar.cc/300" alt="User Avatar" />
                </div>
              </div>
              <span className="ml-3">{"Some Username"}</span>
            </div>

            <ul tabIndex={0} className="dropdown-content menu bg-base-100 rounded-box z-50 w-52 p-2 shadow-md">
              <li>
                <Link to={routes.settings}>Settings</Link>
              </li>
              <li>
                <button onClick={() => navigate(routes.logout)}>Logout</button>
              </li>
            </ul>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 p-6 overflow-auto">{children}</div>
    </div>
  );
};

export default Layout;
