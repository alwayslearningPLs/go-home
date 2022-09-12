import { useState } from 'react';
import ShoppingListModal from './ShoppingListModal';

function App() {
 const [_, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  return (
    <ShoppingListModal show={true} handleClose={handleClose} list={[]}/>
  );
}

export default App;
