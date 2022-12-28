
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { ListingComponent } from './listing/listing.component';
import { ArchiveComponent } from './mail-app/archive/archive.component';
import { DraftsComponent } from './mail-app/drafts/drafts.component';
import { InboxComponent } from './mail-app/inbox/inbox.component';
import { MailAppComponent } from './mail-app/mail-app.component';
import { MailHeaderComponent } from './mail-app/mail-header/mail-header.component';
import { SentComponent } from './mail-app/sent/sent.component';
import { SpamComponent } from './mail-app/spam/spam.component';
import { TrashComponent } from './mail-app/trash/trash.component';
import { RealEstateComponent } from './real-estate/real-estate.component';

const routes: Routes = [
  
  {path:'', redirectTo:'listing', pathMatch:'full'},
  
  {path:'home', component:HomeComponent},
  {path:'listing', component:ListingComponent},
  {path:'real-estate',component:RealEstateComponent},
  {path:'mail-app',component:MailAppComponent},
  {path:'mail-app/inbox',component:InboxComponent},
  {path:'mail-app/sent',component:SentComponent},
  {path:'mail-app/drafts',component:DraftsComponent},
  {path:'mail-app/spam',component:SpamComponent},
  {path:'mail-app/trash',component:TrashComponent},
  {path:'mail-app/archive',component:ArchiveComponent},
  {path:'mail-app/header',component:MailHeaderComponent},
  
  {path:'**', component:ListingComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
