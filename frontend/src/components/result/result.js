import React from 'react';


export class Result extends React.Component{

    constructor(props){
        super(props);
        this.tags=[];
        this.score=this.props.score;
        this.resultData=this.props.data;
        console.log('data in result', this.resultData);
        for(let i=0;i<this.resultData.length;i++){
            if(!this.resultData[i].isCurrectlyAnswered){

                for(let j=0;j<this.resultData[i].tags.length;j++){
                    this.tags.push(this.resultData[i].tags[j]);
                }
            }
        }

        console.log('areas to work on :-', this.tags);
    }

    render(){

        return (
            <div className="mt-3">
            <table className="table">
  <thead className="thead-dark">
    <tr>
      <th scope="col">QUESTION NO</th>
      <th scope="col">RIGHT ANS</th>
      <th scope="col">YOUR ANS</th>
      <th scope="col">MARKS AWARDED</th>

    </tr>
  </thead>
  <tbody>

    {
        this.resultData.map(item=>{
           return <tr>
      <th scope="row">{item.questionId}</th>
      <td>{item.rightAns}</td>
      <td>{item.yourAns}</td>
      <td>{item.marksAwarded}</td>
    </tr>
        })
    }

  </tbody>
</table>

<div className="alert alert-success"> TOTAL SCORE :  {this.score}</div>
<div className="alert alert-danger">Areas you need to work on</div>
<div>
<ul class="list-group">
{
    this.tags.map(tag=>{
        return <li class="list-group-item">{tag}</li>
    })
}

</ul>
</div>
</div>
        )
    }
}