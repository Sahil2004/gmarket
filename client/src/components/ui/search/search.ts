import { Component, Input, OnInit } from '@angular/core';
import { MatInputModule } from '@angular/material/input';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { map, Observable, startWith } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'search',
  templateUrl: 'search.html',
  imports: [
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatButtonModule,
    MatIconModule,
    AsyncPipe,
  ],
})
export class Search implements OnInit {
  @Input() title: string = 'Search';
  @Input() options: string[] = [];
  @Input() symbol: string = 'search';
  @Input() searchHandler: (query: string) => void = (query: string) => {};

  filteredOptions: Observable<string[]> = new Observable<string[]>();

  ngOnInit(): void {
    this.filteredOptions = this.search.valueChanges.pipe(
      startWith(''),
      map((value) => this._filter(value || ''))
    );
  }

  search = new FormControl('');

  private _filter(value: string): string[] {
    const filterVal = value.toLowerCase();
    return this.options.filter((option) => option.toLowerCase().includes(filterVal));
  }
}
