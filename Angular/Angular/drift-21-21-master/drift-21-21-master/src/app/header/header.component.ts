import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  search:string="";
  logo="assets/logo-black.png"

  image="/assets/Image/Sanket.png";

  getsearch(){}
 

}
