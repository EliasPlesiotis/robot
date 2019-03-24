import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';


interface Command {
  Diraction: string;
  Duration;
  Speed;
}

interface File {
  Name: string;
  Commands: Command[];
}

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json'
  })
};


@Component({
  selector: 'app-blocks',
  templateUrl: './blocks.component.html',
  styleUrls: ['./blocks.component.css']
})
export class BlocksComponent implements OnInit {

  moves: Command[] = [];
  files: string[] = [];
  value: string = 'new';
  modified: boolean = false;
  data;
  dur;
  spedd;

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.data = this.http.get<Command[]>(
      'http://192.168.1.6:3000/file/{new}',
       httpOptions)
       .subscribe((res) => {
        this.moves = res;
        console.log(res);
       });

    this.data = this.http.get<string[]>(
      'http://192.168.1.6:3000/files',
       httpOptions)
       .subscribe((res) => {
         this.files = res;
        });
  }

  move(dir: string, speed: number, dur: number) {
    const m: Command = {Diraction: dir, Duration: dur, Speed: speed};
    this.moves.push(m);
    console.log(this.moves);
    this.modified = true;
  }


  delete(m: Command) {
    let i = this.moves.indexOf(m, 0);
    this.moves.splice(i, 1);
    if (this.moves == null) {
      this.moves = [];
    }
    this.modified = true;
  }

  start() {
    alert('saving file ' + this.value);
    this.http.post('http://192.168.1.6:5000/start/' + this.value, null).subscribe((res) => {console.log(res)});
    this.modified = false;
  }

  stop() {
    this.http.post('http://192.168.1.6:5000/stop', null).subscribe((res) => {console.log(res)});
  }

  save(name) {
    let c: File = {Name: name, Commands: this.moves};
    let headers = httpOptions;
    this.http.post<File>(
      'http://192.168.1.6:3000/file/{' + name + '}', c, headers).
      subscribe((res) => {console.log(res)});

    this.data = this.http.get<string[]>(
      'http://192.168.1.6:3000/files',
       httpOptions)
       .subscribe((res) => {
         this.files = res;
        });

    this.modified = false;
  }

  getFile(name) {
    this.data = this.http.get<Command[]>(
      'http://192.168.1.6:3000/file/{' + name + '}',
       httpOptions)
       .subscribe((res) => {
        this.moves = res;
        this.value = name;
        if (this.moves == null) {
          this.moves = []
        }
        console.log(this.moves);
    });
    this.modified = false;
  }

  deleteFile(name) {
    alert('Are you sure')
    this.data = this.http.delete<Command[]>(
      'http://192.168.1.6:3000/file/{' + name + '}',
       httpOptions)
       .subscribe((res) => {
        this.value = name;
        if (this.moves == null) {
          this.moves = []
        }
        console.log(this.moves);
    });

    this.data = this.http.get<string[]>(
      'http://192.168.1.6:3000/files',
       httpOptions)
       .subscribe((res) => {
         this.files = res;
        });

  }

}
