import 'bootstrap/dist/css/bootstrap.min.css';
import {Container} from "react-bootstrap";
import AppNavigation from "./modules/nav.tsx";
import CountriesGrid from "./modules/countries.tsx";

export type SwiftCodeResponse = {
  address: string;
  bankName: string;
  countryISO2: string;
  isHeadquarter: boolean;
  swiftCode: string;
}

function App() {
  return (
    <Container className={"p-2"}>
      <AppNavigation/>
      <Container className="p-5 mb-4 bg-light rounded-3">
        <h2>Select a country</h2>
        <CountriesGrid/>
      </Container>
    </Container>
  )
}

export default App
