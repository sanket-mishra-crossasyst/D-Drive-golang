import { NgModule,  } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';

import { HeaderComponent } from './header/header.component';
import { SidenavComponent } from './sidenav/sidenav.component';
import { HomeComponent } from './home/home.component';

// Angular-Material------>
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatMenuModule }from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import {MatFormFieldModule} from '@angular/material/form-field';  
import {MatSelectModule} from '@angular/material/select';
import {MatCardModule} from '@angular/material/card';
import { DealsClosedGraphComponent } from './listing/deals-closed-graph/deals-closed-graph.component';
import { DashboardListingComponent } from './listing/dashboard-listing/dashboard-listing.component';
import { CurrentPlanComponent } from './listing/current-plan/current-plan.component';
import { ModernHomeComponent } from './listing/modern-home/modern-home.component';

import { RecentActivitiesComponent } from './listing/recent-activities/recent-activities.component';
import { DealersComponent } from './listing/dealers/dealers.component';
import {MatTabsModule} from '@angular/material/tabs';
import { ListingComponent } from './listing/listing.component';
import { PropertiesComponent } from './properties/properties.component';
import {MatExpansionModule} from '@angular/material/expansion';
import { RealEstateComponent } from './real-estate/real-estate.component';
import { MailAppComponent } from './mail-app/mail-app.component';
import { RouterModule } from '@angular/router';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatCheckboxModule} from '@angular/material/checkbox';
import { InboxComponent } from './mail-app/inbox/inbox.component';
import { SentComponent } from './mail-app/sent/sent.component';
import { DraftsComponent } from './mail-app/drafts/drafts.component';
import { SpamComponent } from './mail-app/spam/spam.component';
import { TrashComponent } from './mail-app/trash/trash.component';
import { ArchiveComponent } from './mail-app/archive/archive.component';
import { MailHeaderComponent } from './mail-app/mail-header/mail-header.component';
import { MailSidenavComponent } from './mail-app/mail-sidenav/mail-sidenav.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    SidenavComponent,
    HomeComponent,
    ListingComponent,
    DealsClosedGraphComponent,
    DashboardListingComponent,
    CurrentPlanComponent,
    ModernHomeComponent,
    PropertiesComponent,
    RecentActivitiesComponent,
    DealersComponent,
    RealEstateComponent,
    MailAppComponent,
    InboxComponent,
    SentComponent,
    DraftsComponent,
    SpamComponent,
    TrashComponent,
    ArchiveComponent,
    MailAppComponent,
    MailHeaderComponent,
    MailHeaderComponent,
    MailSidenavComponent,

  
   
   
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FormsModule,
    RouterModule,  
  // Angular-Material------>  
    MatToolbarModule,
    MatMenuModule,
    MatIconModule,
    MatDividerModule,
    MatListModule,
    MatSidenavModule,
    MatIconModule,
    MatFormFieldModule,
    MatSelectModule,
    MatCardModule,
    MatTabsModule,
    MatExpansionModule,
    MatGridListModule,
    MatCheckboxModule,
    MatSidenavModule,

    

    
    
    
  
  
   

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
