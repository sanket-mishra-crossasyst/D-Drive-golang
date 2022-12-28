import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dealers',
  templateUrl: './dealers.component.html',
  styleUrls: ['./dealers.component.scss']
})
export class DealersComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  dealers=[
{
'image':'/assets/Image/noah-owens.jpg',
'name':'Noah Owens',
'rating':'4.9',
'deals':'27 Deals'
},
{
'image':'/assets/Image/ellen-manning.jpg',
'name':'Ellen Manning',
'rating':'4.5',
'deals':'21 Deals'
},
{
'image':'/assets/Image/jones.jpg',
'name':'Angel Jones',
'rating':'4.7',
'deals':'13 Deals'
},
{
'image':'/assets/Image/mila-alba.jpg',
'name':'Dollie Norton',
'rating':'4.4',
'deals':'21 Deals'
},
{
'image':'/assets/Image/steve-smith.jpg',
'name':'Joshua Brian',
'rating':'4.3',
'deals':'07 Deals'
},
{
'image':'/assets/Image/amay.jpg',
'name':'Amay Mathur',
'rating':'4.3',
'deals':'07 Deals'
},
{
'image':'/assets/Image/selena.jpg',
'name':'Selena',
'rating':'4.4',
'deals':'07 Deals'
},
{
'image':'/assets/Image/mathew.jpg',
'name':'Chris Mathew',
'rating':'3.3',
'deals':'07 Deals'
},
{
'image':'/assets/Image/nikki.jpg',
'name':'Nikki Johnson',
'rating':'4.9',
'deals':'07 Deals'
},
{
'image':'',
'name':'Lara Crpft',
'rating':'4.5',
'deals':'07 Deals'
},

  ]

}
