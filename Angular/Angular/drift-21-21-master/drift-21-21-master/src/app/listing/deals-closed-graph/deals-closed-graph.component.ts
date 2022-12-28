import { Component, OnInit } from '@angular/core';
import{Chart, registerables} from 'node_modules/chart.js'
Chart.register(...registerables);


@Component({
  selector: 'app-deals-closed-graph',
  templateUrl: './deals-closed-graph.component.html',
  styleUrls: ['./deals-closed-graph.component.scss']
})
export class DealsClosedGraphComponent implements OnInit {

  constructor() { }

  ngOnInit(): void { this.RenderChart();
  }
  

  RenderChart(){

    const config =new Chart("dealsClosed", {
      type: 'line',
      data:  {
        labels: ['JAN','FEB','MAR','APR','MAY','JUN','JUL','AUG','SEP','OCT','NOV','DEC'],
        datasets: [
          {
          label: 'Queries',
          data: [75,80,85,85,80,70,65,68,72,80,80,70],
          fill: true,
          borderColor: 'rgb(86,68,174)',
        },
        {
          label: 'Closed Deals',
          data: [62,68,73,76,79,76,70,62,59,62,61,56],
          fill: true,
          borderColor: 'rgb(243,72,133)',
        }
      ]
      },
      options: {
        animations: {
          tension: {
            duration: 1000,
            easing: 'linear',
            from: 0.5,
            to: 0,
            loop: true
          }
        },
        scales: {
          y: { // defining min and max so hiding the dataset does not change scale range
            min: 0,
            max: 100,
          }
        }
      }
    });
    



  }
}
