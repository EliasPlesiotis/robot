import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-controller',
  templateUrl: './controller.component.html',
  styleUrls: ['./controller.component.css']
})
export class ControllerComponent implements OnInit {

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  move(dir: string, speed: number) {
    this.http.post('http://192.168.1.6:5000/controller/' + dir + '+' + speed, null).subscribe((res) => {console.log(res)});
  }

  stop() {
    this.http.post('http://192.168.1.6:5000/stop', null).subscribe((res) => {console.log(res)});
  }

}
