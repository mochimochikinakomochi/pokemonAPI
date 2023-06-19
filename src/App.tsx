import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import { Home } from './pages/Home'
import { CreateAccountPage } from './pages/CreateAccountPage'
import { PokemonInfo } from './pages/PokemonInfo'
import { AiFillHome } from 'react-icons/ai'

export const App = () => {
  return (
    <BrowserRouter>
      <div>
        <div className='bg-normal bg-fire bg-water bg-grass bg-electric bg-ice bg-fighting bg-poison bg-ground bg-flying bg-psychic bg-bug bg-rock bg-ghost bg-dragon bg-dark bg-steel bg-fairy'>
          <Link to='/' className="">
            <AiFillHome size={50} color={'#FFF'}/>
          </Link>
          <Link to='/setUser' className="">
            <p>Create User</p>
            <AiFillHome size={50} color={'#FFF'}/>
          </Link>
        </div>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/setUser' element={<CreateAccountPage />} />
          <Route path='/pokemonInfo' element={<PokemonInfo/>} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}
