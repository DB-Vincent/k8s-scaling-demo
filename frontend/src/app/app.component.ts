import { CommonModule } from "@angular/common";
import { ApiService } from "./api.service";
import { Component, OnInit, OnDestroy } from "@angular/core";
import { Subscription, interval } from 'rxjs';

@Component({
  selector: "app-root",
  standalone: true,
  imports: [CommonModule],
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent implements OnInit, OnDestroy {
  replicas: any[] = [];
  error: string | null = null;
  private subscription: Subscription | undefined;

  constructor(private apiService: ApiService) {}

  ngOnInit() {
    this.fetchData();

    this.subscription = interval(5000).subscribe(() => {
      this.fetchData();
    });
  }

  ngOnDestroy() {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  private fetchData() {
    this.error = null; // Reset error before new fetch
    this.apiService.getData().subscribe(
      (response) => {
        this.replicas = response.replicas.map((replica: any) => ({
          ...replica,
          timeSince: this.calculateTimeSince(replica.startTime),
        }));
      },
      (error) => {
        console.error("Error fetching data:", error);
        if (error.status === 500 && error.error.error) {
          this.error = error.error.error;
        } else {
          this.error = "An error occurred while fetching the data";
        }
        this.replicas = [];
      }
    );
  }

  private calculateTimeSince(dateString: string): string {
    const cleanDateString = dateString.replace(/ [A-Z]{3,4}$/, "");

    const startDate = new Date(cleanDateString);
    const now = new Date();

    if (isNaN(startDate.getTime())) {
      return "Invalid date";
    }

    const seconds = Math.floor((now.getTime() - startDate.getTime()) / 1000);
    let interval = seconds / 31536000; // Seconds in a year

    if (interval > 1) {
      return (
        Math.floor(interval) +
        " year" +
        (Math.floor(interval) > 1 ? "s" : "") +
        " ago"
      );
    }
    interval = seconds / 2592000; // Seconds in a month
    if (interval > 1) {
      return (
        Math.floor(interval) +
        " month" +
        (Math.floor(interval) > 1 ? "s" : "") +
        " ago"
      );
    }
    interval = seconds / 86400; // Seconds in a day
    if (interval > 1) {
      return (
        Math.floor(interval) +
        " day" +
        (Math.floor(interval) > 1 ? "s" : "") +
        " ago"
      );
    }
    interval = seconds / 3600; // Seconds in an hour
    if (interval > 1) {
      return (
        Math.floor(interval) +
        " hour" +
        (Math.floor(interval) > 1 ? "s" : "") +
        " ago"
      );
    }
    interval = seconds / 60; // Seconds in a minute
    if (interval > 1) {
      return (
        Math.floor(interval) +
        " minute" +
        (Math.floor(interval) > 1 ? "s" : "") +
        " ago"
      );
    }
    return seconds + " second" + (seconds > 1 ? "s" : "") + " ago";
  }
}
