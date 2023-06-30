import { BrowserRouter, Routes, Route} from 'react-router-dom'
import { ScrollToTop } from './components/ScrollToTop'
import { Home } from './pages/Home'
import { Sidebar } from './components/Sidebar'
import { LoginPage } from './pages/LoginPage'
import { PokemonList } from './pages/PokemonList'
import { PokemonInfo } from './pages/PokemonInfo'
import { MoveList } from './pages/MoveList'
import { MoveInfo } from './pages/MoveInfo'

export const App = () => {
  return (
    <BrowserRouter>
      <ScrollToTop />
      <div>
          <div className='bg-fire bg-water bg-grass bg-electric bg-ice bg-fighting bg-poison bg-ground bg-flying bg-psychic bg-bug bg-rock bg-ghost bg-dragon bg-dark bg-steel bg-fairy bg-all bg-normal'></div>
        <Sidebar />
        <Routes>  
          <Route path='/' element={<Home />} />
          <Route path='/Login' element={<LoginPage />} />
          <Route path='/pokemonList' element={<PokemonList />} />
          <Route path='/pokemonInfo/' element={<PokemonInfo />} />
          <Route path='/moveList' element={<MoveList />} />
          <Route path='/moveInfo/' element={<MoveInfo />} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}
