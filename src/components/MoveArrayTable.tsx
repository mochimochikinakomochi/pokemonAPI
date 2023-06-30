import { Move } from '../domains/move/entity'
import { Link } from 'react-router-dom'

export const MoveArrayTable = ({moves}: {moves: Move[]}) => {

    return (
        <div className='grid grid-cols-3 gap-3'>
        {moves?.map(
            (move) => {
                return(
                    <Link to={`../moveInfo/?moveID=${move?.ID}`}>
                        <div className="grid-item border pl-2 bg-gray-200 hover:bg-gray-300 rounded-[16px] shadow-md w-68">
                            <p className="text-xl text-red-800 pt-5 font-semibold">{move?.ID}</p>
                            <p className="text-2xl font-bold border-b-4 border-red-500 w-fit">{move?.Name}</p>
                            <p className={`bg-${move?.Type} text-white rounded w-fit m-1 px-1`}>{move?.Type}</p>
                            <p className='text-xl'>分類 : {move?.Class}</p>
                            <p className='text-xl'>命中率 :&nbsp;{move?.Accuracy === 0 ? "必中" : move?.Accuracy}</p>
                            <p className='text-xl'>PP :&nbsp;{move?.PP}</p>
                            <p className='text-xl'>威力 :&nbsp;{move?.Class === "変化" ? " ー" : move?.Power}</p>
                            <p className='text-xl'>優先度 :&nbsp;{move?.Priority}</p>
                            <b></b>
                        </div>
                    </Link>
                
                )
            }
        )}
        </div>
    )
}