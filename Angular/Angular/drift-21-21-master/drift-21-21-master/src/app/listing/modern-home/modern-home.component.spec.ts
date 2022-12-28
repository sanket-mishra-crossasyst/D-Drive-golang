import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModernHomeComponent } from './modern-home.component';

describe('ModernHomeComponent', () => {
  let component: ModernHomeComponent;
  let fixture: ComponentFixture<ModernHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ModernHomeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ModernHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
