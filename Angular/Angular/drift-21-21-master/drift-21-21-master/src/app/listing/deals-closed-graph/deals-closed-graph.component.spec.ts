import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DealsClosedGraphComponent } from './deals-closed-graph.component';

describe('DealsClosedGraphComponent', () => {
  let component: DealsClosedGraphComponent;
  let fixture: ComponentFixture<DealsClosedGraphComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DealsClosedGraphComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DealsClosedGraphComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
