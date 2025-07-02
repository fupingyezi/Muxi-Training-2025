import { createBrowserRouter } from "react-router-dom";
import CounterApp from "../pages/Counter/index";
import ThemeApp from "../pages/Theme/index";
import TodoApp from "../pages/Todolist";

const router = createBrowserRouter([
  {
    path: "/",
    element: <CounterApp />,
  },
  {
    path: "/theme",
    element: <ThemeApp />,
  },
  {
    path: "/todo",
    element: <TodoApp />,
  },
]);

export default router;
