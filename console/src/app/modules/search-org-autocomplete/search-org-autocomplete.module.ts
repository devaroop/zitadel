import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatLegacyAutocompleteModule as MatAutocompleteModule } from '@angular/material/legacy-autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatChipsModule } from '@angular/material/chips';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatLegacySelectModule as MatSelectModule } from '@angular/material/legacy-select';
import { TranslateModule } from '@ngx-translate/core';
import { InputModule } from 'src/app/modules/input/input.module';

import { SearchOrgAutocompleteComponent } from './search-org-autocomplete.component';

@NgModule({
  declarations: [SearchOrgAutocompleteComponent],
  imports: [
    CommonModule,
    MatAutocompleteModule,
    MatChipsModule,
    MatButtonModule,
    InputModule,
    MatIconModule,
    ReactiveFormsModule,
    MatProgressSpinnerModule,
    FormsModule,
    TranslateModule,
    MatSelectModule,
  ],
  exports: [SearchOrgAutocompleteComponent],
})
export class SearchOrgAutocompleteModule {}
