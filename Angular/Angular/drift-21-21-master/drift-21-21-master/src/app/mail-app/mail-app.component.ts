import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-mail-app',
  templateUrl: './mail-app.component.html',
  styleUrls: ['./mail-app.component.scss']
})
export class MailAppComponent implements OnInit {
  showAll= true;
  showNone= true;
  values=[
    {
      'image':'',
      'name':'Domnic Harris',
      'discription':'Fusce a libero pellentesque',
      'details':'Maecenas sem arcu, scelerisque in odio vel, porttitor dignissim purus. Sed vehicula commodo porta.'
    },
    {
      'image':'',
      'name':'Garry Sobars',
      'discription':'Nullam id ex at augue pharetra vestibulum eget id mauris.',
      'details':'Cras bibendum tortor tortor, eu luctus risus gravida ut. Suspendisse nisi tortor'
    },
    {
      'image':'',
      'name':'Stella Brown',
      'discription':'Vivamus venenatis tempus ipsum, id finibus libero aliquet convallis.',
      'details':'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce lorem diam, pulvinar id nisl non'
    },
    {
      'image':'',
      'name':'Steve Jonson',
      'discription':'Donec ut ante tristique, gravida justo vitae',
      'details':'dictum at ligula vitae, posuere sagittis augue. Nam vitae eros quis felis consectetur egestas vitae vitae massa.'
    },
    {
      'image':'',
      'name':'Alex Dolgove',
      'discription':'Ut tincidunt massa non elementum fermentum..',
      'details':'Nullam vel ipsum eget odio viverra pellentesque. Nulla auctor eu felis eget vulputate.'
    },
    {
      'image':'',
      'name':'Domnic Brown',
      'discription':'Fusce eu rutrum orci. Aliquam hendrerit metus sit amet interdum',
      'details':'Phasellus ligula purus, eleifend vel massa ut, interdum pulvinar sapien. Nullam a ex nec elit condimentum mattis.'
    },
    {
      'image':'',
      'name':'Brian Lara',
      'discription':'Brian Lara',
      'details':'Aenean facilisis, lorem eget interdum tristique, velit odio tempus orci, sed molestie felis ipsum dignissim leo. '
    }
  ]
  constructor() { }

  ngOnInit(): void {
  }

}
