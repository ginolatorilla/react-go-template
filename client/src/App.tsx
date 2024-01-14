import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="flex flex-col grow m-auto items-center space-y-10">
      <div className="flex">
        <a href="https://vitejs.dev" target="_blank">
          <img
            src={viteLogo}
            className="h-36 p-6  hover:drop-shadow-xl transition-all duration-300"
            alt="Vite logo"
          />
        </a>
        <a href="https://react.dev" target="_blank">
          <img
            src={reactLogo}
            className="h-36 p-6 hover:drop-shadow-xl transition-all animate-[spin_20s_linear_infinite] "
            alt="React logo"
          />
        </a>
      </div>
      <h1 className="text-5xl font-sans">Vite + React</h1>
      <div className="p-8 flex flex-col">
        <button
          className="rounded-lg p-2 font-medium bg-gray-300 cursor-pointer border-solid border-2 hover:border-blue-600 transition-all duration-200"
          onClick={() => setCount((count) => count + 1)}
        >
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="text-gray-800">
        Click on the Vite and React logos to learn more
      </p>
    </div>
  );
}

export default App;
