import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dashboard-listing',
  templateUrl: './dashboard-listing.component.html',
  styleUrls: ['./dashboard-listing.component.scss']
})
export class DashboardListingComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }


  cards=[
    {'name':'LISTED PROPERTY',
     'number':26873,
     'percentage':'03%',
     'status':'This week',
     'width':70,
     'color':'rgb(86,68,174)'
    },
     {'name':'SERVING CITIES',
     'number':384,
     'percentage':'',
     'status':'7 New cities this week',
     'width':50,
     'color':'rgb(233,138,42)'
    },
     {'name':'ONLINE VISITORS',
     'number':84729,
     'percentage':'',
     'status':'Avg. 327 visits daily',
     'width':80,
     'color':'rgb(75,140,40)'
    },
     {'name':'ONLINE QUERIES',
     'number':87239,
     'percentage':'38%',
     'status':'past month',
     'width':90,
     'color':'rgb(243,72,133)'
    },

]

}
