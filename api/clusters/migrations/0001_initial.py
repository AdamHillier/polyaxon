# -*- coding: utf-8 -*-
# Generated by Django 1.11.7 on 2017-12-01 18:50
from __future__ import unicode_literals

from django.conf import settings
import django.contrib.postgres.fields.jsonb
from django.db import migrations, models
import django.db.models.deletion
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name='Cluster',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created_at', models.DateTimeField(auto_now_add=True, db_index=True)),
                ('updated_at', models.DateTimeField(auto_now=True)),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False, unique=True)),
                ('version_api', django.contrib.postgres.fields.jsonb.JSONField(help_text='The cluster version api infos')),
                ('user', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='clusters', to=settings.AUTH_USER_MODEL)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='ClusterErrors',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created_at', models.DateTimeField()),
                ('data', django.contrib.postgres.fields.jsonb.JSONField()),
                ('meta', django.contrib.postgres.fields.jsonb.JSONField()),
                ('level', models.CharField(max_length=16)),
                ('cluster', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='errors', to='clusters.Cluster')),
            ],
        ),
        migrations.CreateModel(
            name='ClusterNode',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False, unique=True)),
                ('name', models.CharField(blank=True, help_text='Name of the node', max_length=256, null=True)),
                ('hostname', models.CharField(blank=True, max_length=256, null=True)),
                ('role', models.CharField(choices=[('Master', 'Master'), ('Worker', 'Worker')], help_text='The role of the node', max_length=6)),
                ('docker_version', models.CharField(blank=True, max_length=128, null=True)),
                ('kubelet_version', models.CharField(max_length=10)),
                ('os_image', models.CharField(max_length=128)),
                ('kernel_version', models.CharField(max_length=128)),
                ('schedulable_taints', models.BooleanField(default=False)),
                ('schedulable_state', models.BooleanField(default=False)),
                ('memory', models.BigIntegerField()),
                ('n_cpus', models.SmallIntegerField()),
                ('n_gpus', models.SmallIntegerField()),
                ('status', models.CharField(choices=[('UNKNOWN', 'UNKNOWN'), ('Ready', 'Ready'), ('NotReady', 'NotReady'), ('Deleted', 'Deleted')], default='UNKNOWN', max_length=24)),
                ('is_current', models.BooleanField(default=True)),
                ('cluster', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='nodes', to='clusters.Cluster')),
            ],
        ),
        migrations.CreateModel(
            name='GPU',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created_at', models.DateTimeField(auto_now_add=True, db_index=True)),
                ('updated_at', models.DateTimeField(auto_now=True)),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False, unique=True)),
                ('serial', models.CharField(max_length=256)),
                ('name', models.CharField(max_length=256)),
                ('device', models.CharField(max_length=256)),
                ('memory', models.BigIntegerField()),
                ('cluster_node', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='gpus', to='clusters.ClusterNode')),
            ],
            options={
                'abstract': False,
            },
        ),
    ]
