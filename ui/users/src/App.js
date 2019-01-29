import React, { Component } from 'react';
import './App.css';
import logo from "./car.svg";
import './bootstrap/css/bootstrap.css';
import CarsTable from './component/CarsTable';
import CommentForm from './component/CommentForm';
import CommentList from './component/CommentList';



class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      comments: [],
      loading: false
    };

    this.addComment = this.addComment.bind(this);
    this.pupulteComments = this.pupulteComments.bind(this);
  }

  pupulteComments(data){
    data.forEach(element => {
      let i = {
        name: element.name,
        message: element.text,
        time: element.CreatedAt
      }
      console.log(i);
        this.setState({
          comments:  [...this.state.comments, i]
        })
      
    })
    this.setState({
      loading: false
    })
  }

  componentDidMount() {
    // loading
    this.setState({ loading: true });

    // get all the comments
    fetch("http://localhost:3100/comments")
      .then(res => res.json())
      .then(this.pupulteComments);
     
  }

   /**
   * Add new comment
   * @param {Object} comment
   */
  addComment(comment) {
    console.log(comment);
    
    /* this.setState({
      loading: false,
      comments: [comment, ...this.state.comments]
    }); */

    this.setState(
      (prevState)=>({
        comments:[comment,...prevState.comments]}),() => {
          console.log(this.state.comments);
      })
  }

  render() {
    return (
      <div className="App">
      <header className="App-header">
      <h1 className="App-title">
            Car Comments
            <span className="px-2" role="img" aria-label="Chat">
              💬
            </span>
          </h1>
          <img src={logo}  alt="logo" />
          <div id="credits">Icons made by <a href="https://www.freepik.com/" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" 			    title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>
        </header>
        
        <CarsTable />

        <div className="row">
          <div className="col-4  pt-3 border-right">
            <h6>Say something about Cars</h6>
            <CommentForm addComment={this.addComment} />
          </div>
          <div className="col-8  pt-3 bg-white">
            <CommentList
              loading={this.state.loading}
              comments={this.state.comments}
            />
          </div>
        </div>
      
       
      </div>
    );
  }
}

export default App;
