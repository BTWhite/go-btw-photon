import { Component, OnInit, Injectable, } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable()
@Component({
  selector: 'app-headers',
  templateUrl: './headers.component.html',
  styleUrls: ['./headers.component.css']
})

export class HeadersComponent implements OnInit {
public id;
public text;

public git="https://api.github.com/search/users"
public local="http://localhost:8886/jsonrpc/"
public data;
public params;
public pre;
public item;
public htmlToAdd;
ngOnInit() {
 
}

  constructor(private _http:HttpClient) {
    
   }

  send(a,b,c){
 
    b= JSON.parse(b)
   
  this.params=JSON.stringify(
    {"id":0,"params":b,"method":a});
 
  this._http.post(c,this.params).subscribe((data:any[]) =>{
 this.data=JSON.stringify(data);

this.pre= JSON.parse(this.data)
this.item=this.pre.result
console.log(Object(this.pre));
this.pre=JSON.stringify(this.pre,null,'\t')
this.item="<tr><td>"+this.item.id +" </td><td>"+this.item.amount / 100000000+" </td><td>"+this.item.senderId +"</td><td>"+this.item.recipientId+"</td><tr>";
this.htmlToAdd=this.item;

  }
)
 







}










}
