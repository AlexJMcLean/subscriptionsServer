import { Links, Meta, Outlet, Scripts } from "@remix-run/react";

const App = () => {
  return (
    <html>
      <head>
        <link rel="icon" href="data:image/x-icon;base64,AA" />
        <Meta />
        <Links />
      </head>
      <body>
        <h1>Hello world</h1>
        <Outlet />
        <Scripts />
      </body>
    </html>
  );
};

export default App;
