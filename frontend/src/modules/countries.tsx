import React from "react";
import {Row, Col, Card} from "react-bootstrap";
import countriesJson from "../../country-list.json";
import {Link} from "react-router";

type Country = {
  emojiFlag: string;
  country: string;
  isoCode: string;
};

const CountriesGrid: React.FC = () => {
  return (
    <Row className="g-3 mt-2">
      {countriesJson.map((country: Country, index: number) => (
        <Col key={index} xs={6} md={3} className="text-center">
          <Row className="w-100">
            <Card as={Link} to={"/country/" + country.isoCode}>
              <Card.Body>{country.emojiFlag} {country.country}</Card.Body>
            </Card>
          </Row>
        </Col>
      ))}
    </Row>
  );
};

export default CountriesGrid;
