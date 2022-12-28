import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-modern-home',
  templateUrl: './modern-home.component.html',
  styleUrls: ['./modern-home.component.scss']
})
export class ModernHomeComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  image="/assets/Image/home-logo.jpg" ;

}
