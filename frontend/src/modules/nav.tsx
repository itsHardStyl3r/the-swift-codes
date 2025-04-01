import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

function AppNavigation() {
  return (
    <Navbar className="bg-body-tertiary mb-2 rounded-3">
      <Container>
        <Navbar.Brand href="/">The Swift Codes</Navbar.Brand>
        <Navbar.Toggle/>
        <Navbar.Collapse id="navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="/">By country</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default AppNavigation;