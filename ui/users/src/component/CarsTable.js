import React, { Component } from 'react';
import ReactTable from "react-table";
import "react-table/react-table.css";

class CarsTable extends React.Component {
	constructor(props) {
  	super(props);
      this.state={items:[]};
      this.pupulteData = this.pupulteData.bind(this);
  }
  pupulteData(data){
    data.forEach(element => {
      let i = {
        model: element.model,
        engine: element.engine,
        brand: element.Manifacturer.name,
        car_type: element.CarType.name,
        fuel_type: element.fuel_type,
        power: element.power

      }
      this.setState({
        items:  [...this.state.items, i]
      })
    })
  }
  componentDidMount(){
  	fetch(`http://localhost:3100/cars`)
     .then(result=>result.json())
    .then(this.pupulteData);
    if (this.state.items.length<=0) {
      console.log("empty")
    }
  }
  render() {
  	return(
        <div className="App">
        <h1 className="my-4"> List of all cars</h1>
        {this.state.items.length ?
            <div>
              <ReactTable className="mx-4 my-4 -striped -highlight" data={this.state.items}
            columns={[
                { Header:"Brand",accessor: "brand"},
                { Header:"Name",accessor:"model"},
                { Header:"Engine",accessor: "engine"},
                { Header:"Car type",accessor: "car_type"},
                { Header:"Fuel type",accessor: "fuel_type"},
                { Header:"Power",accessor: "power"}
            ]} 
            defaultPageSize={10} />
            </div> 
            : <p>Loading...</p>
          }
      </div>
        /*  */
   )
  }
}

export default CarsTable
