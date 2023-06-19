import { BiChevronRight } from 'react-icons/bi'
import { BiChevronLeft } from 'react-icons/bi'


export const PreAndNextButton = ({handleClick}: {handleClick: (isNext: boolean) => void}) => {
    return (
        <div className="flex justify-center">
            <button onClick={() => handleClick(false)} className='bg-gray-700 hover:bg-gray-600 text-white rounded px-6 py-3 m-2'>
                <BiChevronLeft />
            </button>
            <button onClick={() => handleClick(true)} className='bg-gray-700 hover:bg-gray-600 text-white rounded px-6 py-3 m-2'>
                <BiChevronRight />
            </button>
        </div>
    )
}