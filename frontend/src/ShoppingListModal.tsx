import { useEffect, useState } from 'react';
import { Col, Container, Form, InputGroup, Modal, Row, Card } from 'react-bootstrap';
import { BsBagPlus } from 'react-icons/bs';
import FoodUnit from './model/FoodUnit';
import axios from 'axios';

const client = axios.create({
  baseURL: 'http://localhost:8080/food/categories/fruit/units'
})

interface ShoppingListProps {
  show: boolean;
  handleClose: () => void;
  list: FoodUnit[];
}

function renderCard(f: FoodUnit) {
  return (
    <Card className="h-100">
      <Card.Img variant="top" src={f.img} />
      <Card.Body>
        <Card.Title>{f.name}</Card.Title>
        <Card.Text>{f.description}</Card.Text>
      </Card.Body>
    </Card>
  )
}

function renderFoodUnits(list: FoodUnit[]) {
  let e = [];
  for (let i = 0; i < list.length; i += 5) {
    e.push(<Row className="gy-4">{[...Array(Math.min(5, list.length-5))].map((_, j) => <Col className="p-3">{renderCard(list[i+j])}</Col>)}</Row>)
  }
  return <>{e}</>;
}

function ShoppingListModal({ show, handleClose, list }: ShoppingListProps) {
  const [listDisplayed, setListDisplayed] = useState([]);

  useEffect(() => {
    const fetchFoodUnits = async () => {
      let response = await client.get('');
      let data = await response.data;
      setListDisplayed(data)
    }

    fetchFoodUnits();
  }, [])

  return (
    <>
      <Modal show={show} onHide={handleClose} className="modal-xl">
        <Modal.Header closeButton>
          <Modal.Title>Select food units</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Container>
            <Row>
              <Col>
                <Container className="pt-4" fluid="md">
                  <InputGroup>
                    <Form.Control placeholder="Food" aria-label="Food" />
                    <InputGroup.Text id="food-input"><BsBagPlus /></InputGroup.Text>
                  </InputGroup>
                </Container>
              </Col>
            </Row>
            {renderFoodUnits(listDisplayed)}
          </Container>
        </Modal.Body>
      </Modal>
    </>
  );
}

export default ShoppingListModal;