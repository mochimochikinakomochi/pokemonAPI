import { BiChevronRight } from 'react-icons/bi'


export const NextIDButton = ({handleClick}: {handleClick: (isNext: boolean) => void}) => {
    return (
        <button onClick={() => handleClick(true)} className=''>
            <div className='flex items-center bg-white h-24 rounded border-2 border-gray-300'>
                <BiChevronRight size={32}/>
            </div>
        </button>
    )
}