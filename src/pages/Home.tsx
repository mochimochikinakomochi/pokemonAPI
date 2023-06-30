import { Link } from "react-router-dom";

export const Home = () => {
    return (
      <div className="grid grid-cols-2">
        <Link to='/login'>
          <div className="grid-item border pl-2 bg-gray-200 hover:bg-gray-300 rounded-[16px] shadow-md h-48 flex items-center justify-center">
            <div className="text-4xl">さいしょから&nbsp;はじめる</div>
          </div>
        </Link>
        <Link to='/login'>
          <div className="grid-item border pl-2 bg-gray-200 hover:bg-gray-300 rounded-[16px] shadow-md h-48 flex items-center justify-center">
            <div className="text-4xl">つづきから&nbsp;はじめる</div>
          </div>
        </Link>
      </div>
    );
}