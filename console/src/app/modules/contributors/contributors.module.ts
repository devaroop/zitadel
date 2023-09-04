import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatLegacyProgressSpinnerModule as MatProgressSpinnerModule } from '@angular/material/legacy-progress-spinner';
import { MatTooltipModule } from '@angular/material/tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { RoleTransformPipeModule } from 'src/app/pipes/role-transform/role-transform.module';

import { AvatarModule } from '../avatar/avatar.module';
import { ContributorsComponent } from './contributors.component';

@NgModule({
  declarations: [ContributorsComponent],
  imports: [
    CommonModule,
    AvatarModule,
    MatIconModule,
    MatTooltipModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    RoleTransformPipeModule,
    TranslateModule,
  ],
  exports: [ContributorsComponent],
})
export class ContributorsModule {}
