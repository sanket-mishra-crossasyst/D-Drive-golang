import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-recent-activities',
  templateUrl: './recent-activities.component.html',
  styleUrls: ['./recent-activities.component.scss']
})
export class RecentActivitiesComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  

  activities=[
    { 'day':'Today',
    'infoToday':[
  { 'image':'assets/Image/jones.jpg',
  'info':'Mila Alba left a 5 star review on Albama’s House'
  },
  { 'image':'assets/Image/mila-alba.jpg',
  'info':'Callback request from Bob Builder for the property Dimitri House'
  },
  { 'image':'assets/Image/noah-owens.jpg',
  'info':'Congratulations to Tom Moody for joining 10+ club'
  },
  { 'image':'assets/Image/steve-smith.jpg',
  'info':'Norman Dolphi is looking for a house in New Jersy, USA'
  }]
},
{
  'day':'Yesterday',
  'infoToday':[
    { 'image':'assets/Image/jones.jpg',
    'info':'Mila Alba left a 5 star review on Albama’s House'
    },
    { 'image':'assets/Image/mila-alba.jpg',
    'info':'Callback request from Bob Builder for the property Dimitri House'
    },
    { 'image':'assets/Image/noah-owens.jpg',
    'info':'Congratulations to Tom Moody for joining 10+ club'
    },
    { 'image':'assets/Image/steve-smith.jpg',
    'info':'Norman Dolphi is looking for a house in New Jersy, USA'
    }]}
]

}
