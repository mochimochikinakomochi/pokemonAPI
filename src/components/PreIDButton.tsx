import { BiChevronLeft } from 'react-icons/bi'


export const PreIDButton = ({handleClick}: {handleClick: (isNext: boolean) => void}) => {
    return (
        <button onClick={() => handleClick(false)} className=''>
            <div className='flex items-center bg-white h-24 rounded border-2 border-gray-300'>
                <BiChevronLeft size={32}/>
            </div>
        </button>
    )
}