import { Component, OnInit, OnDestroy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from './api.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  replicas: any[] = [];
  private intervalId: any;
  private apiSubscription: Subscription = new Subscription();

  constructor(private apiService: ApiService) {}

  ngOnInit() {
    this.fetchData();

    this.intervalId = setInterval(() => {
      this.fetchData();
    }, 5000);
  }

  private fetchData() {
    const subscription = this.apiService.getData().subscribe(
      response => {
        this.replicas = response.replicas;
        console.log(this.replicas);
      },
      error => {
        console.error('Error fetching data:', error);
      }
    );

    this.apiSubscription.add(subscription);
  }

  ngOnDestroy() {
    if (this.intervalId) {
      clearInterval(this.intervalId);
    }

    this.apiSubscription.unsubscribe();
  }
}
