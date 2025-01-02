import { useEffect, useState } from "react";
import reactLogo from "@assets/react.svg";
import viteLogo from "/vite.svg";
import { hello } from "@utils/server";

function App() {
  const [count, setCount] = useState(0);
  const [helloFromServer, setHelloFromServer] = useState("");

  useEffect(() => {
    hello().then((msg) => {
      setHelloFromServer(msg);
    });
  }, [helloFromServer]);

  return (
    <div className="m-auto flex grow flex-col items-center space-y-10">
      <div className="flex">
        <a href="https://vitejs.dev" target="_blank">
          <img
            src={viteLogo}
            className="h-36 p-6  transition-all duration-300 hover:drop-shadow-xl"
            alt="Vite logo"
          />
        </a>
        <a href="https://react.dev" target="_blank">
          <img
            src={reactLogo}
            className="h-36 animate-[spin_20s_linear_infinite] p-6 transition-all hover:drop-shadow-xl "
            alt="React logo"
          />
        </a>
      </div>
      <h1 className="font-sans text-5xl">Vite + React</h1>
      <div className="flex flex-col p-8">
        <button
          className="cursor-pointer rounded-lg border-2 border-solid bg-gray-300 p-2 font-medium transition-all duration-200 hover:border-blue-600"
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
      <p className="text-gray-800">
        This value is read from the Go server: {helloFromServer}
      </p>
    </div>
  );
}

export default App;
