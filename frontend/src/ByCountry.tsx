import 'bootstrap/dist/css/bootstrap.min.css';
import {Card, Container, ListGroup} from "react-bootstrap";
import AppNavigation from "./modules/nav.tsx";
import {useParams} from "react-router";
import {useEffect, useState} from "react";
import axios from "axios";
import {SwiftCodeResponse} from "./App.tsx";

type Iso2Data = {
  countryISO2: string;
  countryName: string;
  swiftCodes: SwiftCodeResponse[];
};


function ByCountry() {
  const {iso2} = useParams<{ iso2: string }>();

  const [data, setData] = useState<Iso2Data | null>(null);
  useEffect(() => {
    if (iso2) {
      axios.get(`http://localhost:8080/v1/swift-codes/country/${iso2}`).then((response) => {
        setData(response.data);
        console.log(response.data);
      });
    }
  }, [iso2]);

  return (
    <Container className={"p-2"}>
      <AppNavigation/>
      <Container className="p-5 mb-4 bg-light rounded-3">
        <h2>Banks in {data?.countryName}</h2>
        <Card>
          <ListGroup variant="flush">
            {data?.swiftCodes &&
              data.swiftCodes.map((swift, index) => (
                <ListGroup.Item key={index}>{swift.bankName}</ListGroup.Item>
              ))
            }
          </ListGroup>
        </Card>
      </Container>
    </Container>
  )
}

export default ByCountry
