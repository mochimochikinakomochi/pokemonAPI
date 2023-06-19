import { Link } from "react-router-dom";

export const Home = () => {
    return (
      <>
        <div>
          <div className="flex justify-center items-center">
            <div className="w-72 h-72 rounded-full bg-gradient-to-b custom-gradient flex justify-center items-center shadow-2xl shadow-inner inset-O">
              <Link to="/PokemonInfo" className="center hover:bg-gray-200" ></Link>
            </div>
          </div>
          <div className="shadow mx-auto my-16"></div>
        </div>
      </>
    );
}