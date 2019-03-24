import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MenuComponent } from './menu/menu.component';
import { ControllerComponent } from './controller/controller.component';
import { BlocksComponent } from './blocks/blocks.component';
import { CodeComponent } from './code/code.component';
import { FilesComponent } from './files/files.component';


const appRoutes: Routes = [
  { path: 'controller', component: ControllerComponent },
  { path: 'blocks',      component: BlocksComponent },
  { path: 'code', component: CodeComponent },
];

@NgModule({
  declarations: [
    AppComponent,
    MenuComponent,
    ControllerComponent,
    BlocksComponent,
    CodeComponent,
    FilesComponent,
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true }
    ),
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
