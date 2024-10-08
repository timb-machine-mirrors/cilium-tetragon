// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package arm64

// This file was generated by dump-syscalls-info

const (
	SYS_IO_SETUP                = 0
	SYS_IO_DESTROY              = 1
	SYS_IO_SUBMIT               = 2
	SYS_IO_CANCEL               = 3
	SYS_IO_GETEVENTS            = 4
	SYS_SETXATTR                = 5
	SYS_LSETXATTR               = 6
	SYS_FSETXATTR               = 7
	SYS_GETXATTR                = 8
	SYS_LGETXATTR               = 9
	SYS_FGETXATTR               = 10
	SYS_LISTXATTR               = 11
	SYS_LLISTXATTR              = 12
	SYS_FLISTXATTR              = 13
	SYS_REMOVEXATTR             = 14
	SYS_LREMOVEXATTR            = 15
	SYS_FREMOVEXATTR            = 16
	SYS_GETCWD                  = 17
	SYS_LOOKUP_DCOOKIE          = 18
	SYS_EVENTFD2                = 19
	SYS_EPOLL_CREATE1           = 20
	SYS_EPOLL_CTL               = 21
	SYS_EPOLL_PWAIT             = 22
	SYS_DUP                     = 23
	SYS_DUP3                    = 24
	SYS_FCNTL                   = 25
	SYS_INOTIFY_INIT1           = 26
	SYS_INOTIFY_ADD_WATCH       = 27
	SYS_INOTIFY_RM_WATCH        = 28
	SYS_IOCTL                   = 29
	SYS_IOPRIO_SET              = 30
	SYS_IOPRIO_GET              = 31
	SYS_FLOCK                   = 32
	SYS_MKNODAT                 = 33
	SYS_MKDIRAT                 = 34
	SYS_UNLINKAT                = 35
	SYS_SYMLINKAT               = 36
	SYS_LINKAT                  = 37
	SYS_RENAMEAT                = 38
	SYS_UMOUNT2                 = 39
	SYS_MOUNT                   = 40
	SYS_PIVOT_ROOT              = 41
	SYS_NFSSERVCTL              = 42
	SYS_STATFS                  = 43
	SYS_FSTATFS                 = 44
	SYS_TRUNCATE                = 45
	SYS_FTRUNCATE               = 46
	SYS_FALLOCATE               = 47
	SYS_FACCESSAT               = 48
	SYS_CHDIR                   = 49
	SYS_FCHDIR                  = 50
	SYS_CHROOT                  = 51
	SYS_FCHMOD                  = 52
	SYS_FCHMODAT                = 53
	SYS_FCHOWNAT                = 54
	SYS_FCHOWN                  = 55
	SYS_OPENAT                  = 56
	SYS_CLOSE                   = 57
	SYS_VHANGUP                 = 58
	SYS_PIPE2                   = 59
	SYS_QUOTACTL                = 60
	SYS_GETDENTS64              = 61
	SYS_LSEEK                   = 62
	SYS_READ                    = 63
	SYS_WRITE                   = 64
	SYS_READV                   = 65
	SYS_WRITEV                  = 66
	SYS_PREAD64                 = 67
	SYS_PWRITE64                = 68
	SYS_PREADV                  = 69
	SYS_PWRITEV                 = 70
	SYS_SENDFILE                = 71
	SYS_PSELECT6                = 72
	SYS_PPOLL                   = 73
	SYS_SIGNALFD4               = 74
	SYS_VMSPLICE                = 75
	SYS_SPLICE                  = 76
	SYS_TEE                     = 77
	SYS_READLINKAT              = 78
	SYS_NEWFSTATAT              = 79
	SYS_FSTAT                   = 80
	SYS_SYNC                    = 81
	SYS_FSYNC                   = 82
	SYS_FDATASYNC               = 83
	SYS_SYNC_FILE_RANGE         = 84
	SYS_TIMERFD_CREATE          = 85
	SYS_TIMERFD_SETTIME         = 86
	SYS_TIMERFD_GETTIME         = 87
	SYS_UTIMENSAT               = 88
	SYS_ACCT                    = 89
	SYS_CAPGET                  = 90
	SYS_CAPSET                  = 91
	SYS_PERSONALITY             = 92
	SYS_EXIT                    = 93
	SYS_EXIT_GROUP              = 94
	SYS_WAITID                  = 95
	SYS_SET_TID_ADDRESS         = 96
	SYS_UNSHARE                 = 97
	SYS_FUTEX                   = 98
	SYS_SET_ROBUST_LIST         = 99
	SYS_GET_ROBUST_LIST         = 100
	SYS_NANOSLEEP               = 101
	SYS_GETITIMER               = 102
	SYS_SETITIMER               = 103
	SYS_KEXEC_LOAD              = 104
	SYS_INIT_MODULE             = 105
	SYS_DELETE_MODULE           = 106
	SYS_TIMER_CREATE            = 107
	SYS_TIMER_GETTIME           = 108
	SYS_TIMER_GETOVERRUN        = 109
	SYS_TIMER_SETTIME           = 110
	SYS_TIMER_DELETE            = 111
	SYS_CLOCK_SETTIME           = 112
	SYS_CLOCK_GETTIME           = 113
	SYS_CLOCK_GETRES            = 114
	SYS_CLOCK_NANOSLEEP         = 115
	SYS_SYSLOG                  = 116
	SYS_PTRACE                  = 117
	SYS_SCHED_SETPARAM          = 118
	SYS_SCHED_SETSCHEDULER      = 119
	SYS_SCHED_GETSCHEDULER      = 120
	SYS_SCHED_GETPARAM          = 121
	SYS_SCHED_SETAFFINITY       = 122
	SYS_SCHED_GETAFFINITY       = 123
	SYS_SCHED_YIELD             = 124
	SYS_SCHED_GET_PRIORITY_MAX  = 125
	SYS_SCHED_GET_PRIORITY_MIN  = 126
	SYS_SCHED_RR_GET_INTERVAL   = 127
	SYS_RESTART_SYSCALL         = 128
	SYS_KILL                    = 129
	SYS_TKILL                   = 130
	SYS_TGKILL                  = 131
	SYS_SIGALTSTACK             = 132
	SYS_RT_SIGSUSPEND           = 133
	SYS_RT_SIGACTION            = 134
	SYS_RT_SIGPROCMASK          = 135
	SYS_RT_SIGPENDING           = 136
	SYS_RT_SIGTIMEDWAIT         = 137
	SYS_RT_SIGQUEUEINFO         = 138
	SYS_RT_SIGRETURN            = 139
	SYS_SETPRIORITY             = 140
	SYS_GETPRIORITY             = 141
	SYS_REBOOT                  = 142
	SYS_SETREGID                = 143
	SYS_SETGID                  = 144
	SYS_SETREUID                = 145
	SYS_SETUID                  = 146
	SYS_SETRESUID               = 147
	SYS_GETRESUID               = 148
	SYS_SETRESGID               = 149
	SYS_GETRESGID               = 150
	SYS_SETFSUID                = 151
	SYS_SETFSGID                = 152
	SYS_TIMES                   = 153
	SYS_SETPGID                 = 154
	SYS_GETPGID                 = 155
	SYS_GETSID                  = 156
	SYS_SETSID                  = 157
	SYS_GETGROUPS               = 158
	SYS_SETGROUPS               = 159
	SYS_UNAME                   = 160
	SYS_SETHOSTNAME             = 161
	SYS_SETDOMAINNAME           = 162
	SYS_GETRLIMIT               = 163
	SYS_SETRLIMIT               = 164
	SYS_GETRUSAGE               = 165
	SYS_UMASK                   = 166
	SYS_PRCTL                   = 167
	SYS_GETCPU                  = 168
	SYS_GETTIMEOFDAY            = 169
	SYS_SETTIMEOFDAY            = 170
	SYS_ADJTIMEX                = 171
	SYS_GETPID                  = 172
	SYS_GETPPID                 = 173
	SYS_GETUID                  = 174
	SYS_GETEUID                 = 175
	SYS_GETGID                  = 176
	SYS_GETEGID                 = 177
	SYS_GETTID                  = 178
	SYS_SYSINFO                 = 179
	SYS_MQ_OPEN                 = 180
	SYS_MQ_UNLINK               = 181
	SYS_MQ_TIMEDSEND            = 182
	SYS_MQ_TIMEDRECEIVE         = 183
	SYS_MQ_NOTIFY               = 184
	SYS_MQ_GETSETATTR           = 185
	SYS_MSGGET                  = 186
	SYS_MSGCTL                  = 187
	SYS_MSGRCV                  = 188
	SYS_MSGSND                  = 189
	SYS_SEMGET                  = 190
	SYS_SEMCTL                  = 191
	SYS_SEMTIMEDOP              = 192
	SYS_SEMOP                   = 193
	SYS_SHMGET                  = 194
	SYS_SHMCTL                  = 195
	SYS_SHMAT                   = 196
	SYS_SHMDT                   = 197
	SYS_SOCKET                  = 198
	SYS_SOCKETPAIR              = 199
	SYS_BIND                    = 200
	SYS_LISTEN                  = 201
	SYS_ACCEPT                  = 202
	SYS_CONNECT                 = 203
	SYS_GETSOCKNAME             = 204
	SYS_GETPEERNAME             = 205
	SYS_SENDTO                  = 206
	SYS_RECVFROM                = 207
	SYS_SETSOCKOPT              = 208
	SYS_GETSOCKOPT              = 209
	SYS_SHUTDOWN                = 210
	SYS_SENDMSG                 = 211
	SYS_RECVMSG                 = 212
	SYS_READAHEAD               = 213
	SYS_BRK                     = 214
	SYS_MUNMAP                  = 215
	SYS_MREMAP                  = 216
	SYS_ADD_KEY                 = 217
	SYS_REQUEST_KEY             = 218
	SYS_KEYCTL                  = 219
	SYS_CLONE                   = 220
	SYS_EXECVE                  = 221
	SYS_MMAP                    = 222
	SYS_FADVISE64               = 223
	SYS_SWAPON                  = 224
	SYS_SWAPOFF                 = 225
	SYS_MPROTECT                = 226
	SYS_MSYNC                   = 227
	SYS_MLOCK                   = 228
	SYS_MUNLOCK                 = 229
	SYS_MLOCKALL                = 230
	SYS_MUNLOCKALL              = 231
	SYS_MINCORE                 = 232
	SYS_MADVISE                 = 233
	SYS_REMAP_FILE_PAGES        = 234
	SYS_MBIND                   = 235
	SYS_GET_MEMPOLICY           = 236
	SYS_SET_MEMPOLICY           = 237
	SYS_MIGRATE_PAGES           = 238
	SYS_MOVE_PAGES              = 239
	SYS_RT_TGSIGQUEUEINFO       = 240
	SYS_PERF_EVENT_OPEN         = 241
	SYS_ACCEPT4                 = 242
	SYS_RECVMMSG                = 243
	SYS_WAIT4                   = 260
	SYS_PRLIMIT64               = 261
	SYS_FANOTIFY_INIT           = 262
	SYS_FANOTIFY_MARK           = 263
	SYS_NAME_TO_HANDLE_AT       = 264
	SYS_OPEN_BY_HANDLE_AT       = 265
	SYS_CLOCK_ADJTIME           = 266
	SYS_SYNCFS                  = 267
	SYS_SETNS                   = 268
	SYS_SENDMMSG                = 269
	SYS_PROCESS_VM_READV        = 270
	SYS_PROCESS_VM_WRITEV       = 271
	SYS_KCMP                    = 272
	SYS_FINIT_MODULE            = 273
	SYS_SCHED_SETATTR           = 274
	SYS_SCHED_GETATTR           = 275
	SYS_RENAMEAT2               = 276
	SYS_SECCOMP                 = 277
	SYS_GETRANDOM               = 278
	SYS_MEMFD_CREATE            = 279
	SYS_BPF                     = 280
	SYS_EXECVEAT                = 281
	SYS_USERFAULTFD             = 282
	SYS_MEMBARRIER              = 283
	SYS_MLOCK2                  = 284
	SYS_COPY_FILE_RANGE         = 285
	SYS_PREADV2                 = 286
	SYS_PWRITEV2                = 287
	SYS_PKEY_MPROTECT           = 288
	SYS_PKEY_ALLOC              = 289
	SYS_PKEY_FREE               = 290
	SYS_STATX                   = 291
	SYS_IO_PGETEVENTS           = 292
	SYS_RSEQ                    = 293
	SYS_KEXEC_FILE_LOAD         = 294
	SYS_PIDFD_SEND_SIGNAL       = 424
	SYS_IO_URING_SETUP          = 425
	SYS_IO_URING_ENTER          = 426
	SYS_IO_URING_REGISTER       = 427
	SYS_OPEN_TREE               = 428
	SYS_MOVE_MOUNT              = 429
	SYS_FSOPEN                  = 430
	SYS_FSCONFIG                = 431
	SYS_FSMOUNT                 = 432
	SYS_FSPICK                  = 433
	SYS_PIDFD_OPEN              = 434
	SYS_CLONE3                  = 435
	SYS_CLOSE_RANGE             = 436
	SYS_OPENAT2                 = 437
	SYS_PIDFD_GETFD             = 438
	SYS_FACCESSAT2              = 439
	SYS_PROCESS_MADVISE         = 440
	SYS_EPOLL_PWAIT2            = 441
	SYS_MOUNT_SETATTR           = 442
	SYS_QUOTACTL_FD             = 443
	SYS_LANDLOCK_CREATE_RULESET = 444
	SYS_LANDLOCK_ADD_RULE       = 445
	SYS_LANDLOCK_RESTRICT_SELF  = 446
	SYS_MEMFD_SECRET            = 447
	SYS_PROCESS_MRELEASE        = 448
	SYS_FUTEX_WAITV             = 449
	SYS_SET_MEMPOLICY_HOME_NODE = 450
	SYS_CACHESTAT               = 451
	SYS_FCHMODAT2               = 452
	SYS_MAP_SHADOW_STACK        = 453
	SYS_FUTEX_WAKE              = 454
	SYS_FUTEX_WAIT              = 455
	SYS_FUTEX_REQUEUE           = 456
	SYS_STATMOUNT               = 457
	SYS_LISTMOUNT               = 458
	SYS_LSM_GET_SELF_ATTR       = 459
	SYS_LSM_SET_SELF_ATTR       = 460
	SYS_LSM_LIST_MODULES        = 461
	SYS_MSEAL                   = 462
)

var Names = map[int]string{
	SYS_IO_SETUP:                "io_setup",
	SYS_IO_DESTROY:              "io_destroy",
	SYS_IO_SUBMIT:               "io_submit",
	SYS_IO_CANCEL:               "io_cancel",
	SYS_IO_GETEVENTS:            "io_getevents",
	SYS_SETXATTR:                "setxattr",
	SYS_LSETXATTR:               "lsetxattr",
	SYS_FSETXATTR:               "fsetxattr",
	SYS_GETXATTR:                "getxattr",
	SYS_LGETXATTR:               "lgetxattr",
	SYS_FGETXATTR:               "fgetxattr",
	SYS_LISTXATTR:               "listxattr",
	SYS_LLISTXATTR:              "llistxattr",
	SYS_FLISTXATTR:              "flistxattr",
	SYS_REMOVEXATTR:             "removexattr",
	SYS_LREMOVEXATTR:            "lremovexattr",
	SYS_FREMOVEXATTR:            "fremovexattr",
	SYS_GETCWD:                  "getcwd",
	SYS_LOOKUP_DCOOKIE:          "lookup_dcookie",
	SYS_EVENTFD2:                "eventfd2",
	SYS_EPOLL_CREATE1:           "epoll_create1",
	SYS_EPOLL_CTL:               "epoll_ctl",
	SYS_EPOLL_PWAIT:             "epoll_pwait",
	SYS_DUP:                     "dup",
	SYS_DUP3:                    "dup3",
	SYS_FCNTL:                   "fcntl",
	SYS_INOTIFY_INIT1:           "inotify_init1",
	SYS_INOTIFY_ADD_WATCH:       "inotify_add_watch",
	SYS_INOTIFY_RM_WATCH:        "inotify_rm_watch",
	SYS_IOCTL:                   "ioctl",
	SYS_IOPRIO_SET:              "ioprio_set",
	SYS_IOPRIO_GET:              "ioprio_get",
	SYS_FLOCK:                   "flock",
	SYS_MKNODAT:                 "mknodat",
	SYS_MKDIRAT:                 "mkdirat",
	SYS_UNLINKAT:                "unlinkat",
	SYS_SYMLINKAT:               "symlinkat",
	SYS_LINKAT:                  "linkat",
	SYS_RENAMEAT:                "renameat",
	SYS_UMOUNT2:                 "umount2",
	SYS_MOUNT:                   "mount",
	SYS_PIVOT_ROOT:              "pivot_root",
	SYS_NFSSERVCTL:              "nfsservctl",
	SYS_STATFS:                  "statfs",
	SYS_FSTATFS:                 "fstatfs",
	SYS_TRUNCATE:                "truncate",
	SYS_FTRUNCATE:               "ftruncate",
	SYS_FALLOCATE:               "fallocate",
	SYS_FACCESSAT:               "faccessat",
	SYS_CHDIR:                   "chdir",
	SYS_FCHDIR:                  "fchdir",
	SYS_CHROOT:                  "chroot",
	SYS_FCHMOD:                  "fchmod",
	SYS_FCHMODAT:                "fchmodat",
	SYS_FCHOWNAT:                "fchownat",
	SYS_FCHOWN:                  "fchown",
	SYS_OPENAT:                  "openat",
	SYS_CLOSE:                   "close",
	SYS_VHANGUP:                 "vhangup",
	SYS_PIPE2:                   "pipe2",
	SYS_QUOTACTL:                "quotactl",
	SYS_GETDENTS64:              "getdents64",
	SYS_LSEEK:                   "lseek",
	SYS_READ:                    "read",
	SYS_WRITE:                   "write",
	SYS_READV:                   "readv",
	SYS_WRITEV:                  "writev",
	SYS_PREAD64:                 "pread64",
	SYS_PWRITE64:                "pwrite64",
	SYS_PREADV:                  "preadv",
	SYS_PWRITEV:                 "pwritev",
	SYS_SENDFILE:                "sendfile",
	SYS_PSELECT6:                "pselect6",
	SYS_PPOLL:                   "ppoll",
	SYS_SIGNALFD4:               "signalfd4",
	SYS_VMSPLICE:                "vmsplice",
	SYS_SPLICE:                  "splice",
	SYS_TEE:                     "tee",
	SYS_READLINKAT:              "readlinkat",
	SYS_NEWFSTATAT:              "newfstatat",
	SYS_FSTAT:                   "fstat",
	SYS_SYNC:                    "sync",
	SYS_FSYNC:                   "fsync",
	SYS_FDATASYNC:               "fdatasync",
	SYS_SYNC_FILE_RANGE:         "sync_file_range",
	SYS_TIMERFD_CREATE:          "timerfd_create",
	SYS_TIMERFD_SETTIME:         "timerfd_settime",
	SYS_TIMERFD_GETTIME:         "timerfd_gettime",
	SYS_UTIMENSAT:               "utimensat",
	SYS_ACCT:                    "acct",
	SYS_CAPGET:                  "capget",
	SYS_CAPSET:                  "capset",
	SYS_PERSONALITY:             "personality",
	SYS_EXIT:                    "exit",
	SYS_EXIT_GROUP:              "exit_group",
	SYS_WAITID:                  "waitid",
	SYS_SET_TID_ADDRESS:         "set_tid_address",
	SYS_UNSHARE:                 "unshare",
	SYS_FUTEX:                   "futex",
	SYS_SET_ROBUST_LIST:         "set_robust_list",
	SYS_GET_ROBUST_LIST:         "get_robust_list",
	SYS_NANOSLEEP:               "nanosleep",
	SYS_GETITIMER:               "getitimer",
	SYS_SETITIMER:               "setitimer",
	SYS_KEXEC_LOAD:              "kexec_load",
	SYS_INIT_MODULE:             "init_module",
	SYS_DELETE_MODULE:           "delete_module",
	SYS_TIMER_CREATE:            "timer_create",
	SYS_TIMER_GETTIME:           "timer_gettime",
	SYS_TIMER_GETOVERRUN:        "timer_getoverrun",
	SYS_TIMER_SETTIME:           "timer_settime",
	SYS_TIMER_DELETE:            "timer_delete",
	SYS_CLOCK_SETTIME:           "clock_settime",
	SYS_CLOCK_GETTIME:           "clock_gettime",
	SYS_CLOCK_GETRES:            "clock_getres",
	SYS_CLOCK_NANOSLEEP:         "clock_nanosleep",
	SYS_SYSLOG:                  "syslog",
	SYS_PTRACE:                  "ptrace",
	SYS_SCHED_SETPARAM:          "sched_setparam",
	SYS_SCHED_SETSCHEDULER:      "sched_setscheduler",
	SYS_SCHED_GETSCHEDULER:      "sched_getscheduler",
	SYS_SCHED_GETPARAM:          "sched_getparam",
	SYS_SCHED_SETAFFINITY:       "sched_setaffinity",
	SYS_SCHED_GETAFFINITY:       "sched_getaffinity",
	SYS_SCHED_YIELD:             "sched_yield",
	SYS_SCHED_GET_PRIORITY_MAX:  "sched_get_priority_max",
	SYS_SCHED_GET_PRIORITY_MIN:  "sched_get_priority_min",
	SYS_SCHED_RR_GET_INTERVAL:   "sched_rr_get_interval",
	SYS_RESTART_SYSCALL:         "restart_syscall",
	SYS_KILL:                    "kill",
	SYS_TKILL:                   "tkill",
	SYS_TGKILL:                  "tgkill",
	SYS_SIGALTSTACK:             "sigaltstack",
	SYS_RT_SIGSUSPEND:           "rt_sigsuspend",
	SYS_RT_SIGACTION:            "rt_sigaction",
	SYS_RT_SIGPROCMASK:          "rt_sigprocmask",
	SYS_RT_SIGPENDING:           "rt_sigpending",
	SYS_RT_SIGTIMEDWAIT:         "rt_sigtimedwait",
	SYS_RT_SIGQUEUEINFO:         "rt_sigqueueinfo",
	SYS_RT_SIGRETURN:            "rt_sigreturn",
	SYS_SETPRIORITY:             "setpriority",
	SYS_GETPRIORITY:             "getpriority",
	SYS_REBOOT:                  "reboot",
	SYS_SETREGID:                "setregid",
	SYS_SETGID:                  "setgid",
	SYS_SETREUID:                "setreuid",
	SYS_SETUID:                  "setuid",
	SYS_SETRESUID:               "setresuid",
	SYS_GETRESUID:               "getresuid",
	SYS_SETRESGID:               "setresgid",
	SYS_GETRESGID:               "getresgid",
	SYS_SETFSUID:                "setfsuid",
	SYS_SETFSGID:                "setfsgid",
	SYS_TIMES:                   "times",
	SYS_SETPGID:                 "setpgid",
	SYS_GETPGID:                 "getpgid",
	SYS_GETSID:                  "getsid",
	SYS_SETSID:                  "setsid",
	SYS_GETGROUPS:               "getgroups",
	SYS_SETGROUPS:               "setgroups",
	SYS_UNAME:                   "uname",
	SYS_SETHOSTNAME:             "sethostname",
	SYS_SETDOMAINNAME:           "setdomainname",
	SYS_GETRLIMIT:               "getrlimit",
	SYS_SETRLIMIT:               "setrlimit",
	SYS_GETRUSAGE:               "getrusage",
	SYS_UMASK:                   "umask",
	SYS_PRCTL:                   "prctl",
	SYS_GETCPU:                  "getcpu",
	SYS_GETTIMEOFDAY:            "gettimeofday",
	SYS_SETTIMEOFDAY:            "settimeofday",
	SYS_ADJTIMEX:                "adjtimex",
	SYS_GETPID:                  "getpid",
	SYS_GETPPID:                 "getppid",
	SYS_GETUID:                  "getuid",
	SYS_GETEUID:                 "geteuid",
	SYS_GETGID:                  "getgid",
	SYS_GETEGID:                 "getegid",
	SYS_GETTID:                  "gettid",
	SYS_SYSINFO:                 "sysinfo",
	SYS_MQ_OPEN:                 "mq_open",
	SYS_MQ_UNLINK:               "mq_unlink",
	SYS_MQ_TIMEDSEND:            "mq_timedsend",
	SYS_MQ_TIMEDRECEIVE:         "mq_timedreceive",
	SYS_MQ_NOTIFY:               "mq_notify",
	SYS_MQ_GETSETATTR:           "mq_getsetattr",
	SYS_MSGGET:                  "msgget",
	SYS_MSGCTL:                  "msgctl",
	SYS_MSGRCV:                  "msgrcv",
	SYS_MSGSND:                  "msgsnd",
	SYS_SEMGET:                  "semget",
	SYS_SEMCTL:                  "semctl",
	SYS_SEMTIMEDOP:              "semtimedop",
	SYS_SEMOP:                   "semop",
	SYS_SHMGET:                  "shmget",
	SYS_SHMCTL:                  "shmctl",
	SYS_SHMAT:                   "shmat",
	SYS_SHMDT:                   "shmdt",
	SYS_SOCKET:                  "socket",
	SYS_SOCKETPAIR:              "socketpair",
	SYS_BIND:                    "bind",
	SYS_LISTEN:                  "listen",
	SYS_ACCEPT:                  "accept",
	SYS_CONNECT:                 "connect",
	SYS_GETSOCKNAME:             "getsockname",
	SYS_GETPEERNAME:             "getpeername",
	SYS_SENDTO:                  "sendto",
	SYS_RECVFROM:                "recvfrom",
	SYS_SETSOCKOPT:              "setsockopt",
	SYS_GETSOCKOPT:              "getsockopt",
	SYS_SHUTDOWN:                "shutdown",
	SYS_SENDMSG:                 "sendmsg",
	SYS_RECVMSG:                 "recvmsg",
	SYS_READAHEAD:               "readahead",
	SYS_BRK:                     "brk",
	SYS_MUNMAP:                  "munmap",
	SYS_MREMAP:                  "mremap",
	SYS_ADD_KEY:                 "add_key",
	SYS_REQUEST_KEY:             "request_key",
	SYS_KEYCTL:                  "keyctl",
	SYS_CLONE:                   "clone",
	SYS_EXECVE:                  "execve",
	SYS_MMAP:                    "mmap",
	SYS_FADVISE64:               "fadvise64",
	SYS_SWAPON:                  "swapon",
	SYS_SWAPOFF:                 "swapoff",
	SYS_MPROTECT:                "mprotect",
	SYS_MSYNC:                   "msync",
	SYS_MLOCK:                   "mlock",
	SYS_MUNLOCK:                 "munlock",
	SYS_MLOCKALL:                "mlockall",
	SYS_MUNLOCKALL:              "munlockall",
	SYS_MINCORE:                 "mincore",
	SYS_MADVISE:                 "madvise",
	SYS_REMAP_FILE_PAGES:        "remap_file_pages",
	SYS_MBIND:                   "mbind",
	SYS_GET_MEMPOLICY:           "get_mempolicy",
	SYS_SET_MEMPOLICY:           "set_mempolicy",
	SYS_MIGRATE_PAGES:           "migrate_pages",
	SYS_MOVE_PAGES:              "move_pages",
	SYS_RT_TGSIGQUEUEINFO:       "rt_tgsigqueueinfo",
	SYS_PERF_EVENT_OPEN:         "perf_event_open",
	SYS_ACCEPT4:                 "accept4",
	SYS_RECVMMSG:                "recvmmsg",
	SYS_WAIT4:                   "wait4",
	SYS_PRLIMIT64:               "prlimit64",
	SYS_FANOTIFY_INIT:           "fanotify_init",
	SYS_FANOTIFY_MARK:           "fanotify_mark",
	SYS_NAME_TO_HANDLE_AT:       "name_to_handle_at",
	SYS_OPEN_BY_HANDLE_AT:       "open_by_handle_at",
	SYS_CLOCK_ADJTIME:           "clock_adjtime",
	SYS_SYNCFS:                  "syncfs",
	SYS_SETNS:                   "setns",
	SYS_SENDMMSG:                "sendmmsg",
	SYS_PROCESS_VM_READV:        "process_vm_readv",
	SYS_PROCESS_VM_WRITEV:       "process_vm_writev",
	SYS_KCMP:                    "kcmp",
	SYS_FINIT_MODULE:            "finit_module",
	SYS_SCHED_SETATTR:           "sched_setattr",
	SYS_SCHED_GETATTR:           "sched_getattr",
	SYS_RENAMEAT2:               "renameat2",
	SYS_SECCOMP:                 "seccomp",
	SYS_GETRANDOM:               "getrandom",
	SYS_MEMFD_CREATE:            "memfd_create",
	SYS_BPF:                     "bpf",
	SYS_EXECVEAT:                "execveat",
	SYS_USERFAULTFD:             "userfaultfd",
	SYS_MEMBARRIER:              "membarrier",
	SYS_MLOCK2:                  "mlock2",
	SYS_COPY_FILE_RANGE:         "copy_file_range",
	SYS_PREADV2:                 "preadv2",
	SYS_PWRITEV2:                "pwritev2",
	SYS_PKEY_MPROTECT:           "pkey_mprotect",
	SYS_PKEY_ALLOC:              "pkey_alloc",
	SYS_PKEY_FREE:               "pkey_free",
	SYS_STATX:                   "statx",
	SYS_IO_PGETEVENTS:           "io_pgetevents",
	SYS_RSEQ:                    "rseq",
	SYS_KEXEC_FILE_LOAD:         "kexec_file_load",
	SYS_PIDFD_SEND_SIGNAL:       "pidfd_send_signal",
	SYS_IO_URING_SETUP:          "io_uring_setup",
	SYS_IO_URING_ENTER:          "io_uring_enter",
	SYS_IO_URING_REGISTER:       "io_uring_register",
	SYS_OPEN_TREE:               "open_tree",
	SYS_MOVE_MOUNT:              "move_mount",
	SYS_FSOPEN:                  "fsopen",
	SYS_FSCONFIG:                "fsconfig",
	SYS_FSMOUNT:                 "fsmount",
	SYS_FSPICK:                  "fspick",
	SYS_PIDFD_OPEN:              "pidfd_open",
	SYS_CLONE3:                  "clone3",
	SYS_CLOSE_RANGE:             "close_range",
	SYS_OPENAT2:                 "openat2",
	SYS_PIDFD_GETFD:             "pidfd_getfd",
	SYS_FACCESSAT2:              "faccessat2",
	SYS_PROCESS_MADVISE:         "process_madvise",
	SYS_EPOLL_PWAIT2:            "epoll_pwait2",
	SYS_MOUNT_SETATTR:           "mount_setattr",
	SYS_QUOTACTL_FD:             "quotactl_fd",
	SYS_LANDLOCK_CREATE_RULESET: "landlock_create_ruleset",
	SYS_LANDLOCK_ADD_RULE:       "landlock_add_rule",
	SYS_LANDLOCK_RESTRICT_SELF:  "landlock_restrict_self",
	SYS_MEMFD_SECRET:            "memfd_secret",
	SYS_PROCESS_MRELEASE:        "process_mrelease",
	SYS_FUTEX_WAITV:             "futex_waitv",
	SYS_SET_MEMPOLICY_HOME_NODE: "set_mempolicy_home_node",
	SYS_CACHESTAT:               "cachestat",
	SYS_FCHMODAT2:               "fchmodat2",
	SYS_MAP_SHADOW_STACK:        "map_shadow_stack",
	SYS_FUTEX_WAKE:              "futex_wake",
	SYS_FUTEX_WAIT:              "futex_wait",
	SYS_FUTEX_REQUEUE:           "futex_requeue",
	SYS_STATMOUNT:               "statmount",
	SYS_LISTMOUNT:               "listmount",
	SYS_LSM_GET_SELF_ATTR:       "lsm_get_self_attr",
	SYS_LSM_SET_SELF_ATTR:       "lsm_set_self_attr",
	SYS_LSM_LIST_MODULES:        "lsm_list_modules",
	SYS_MSEAL:                   "mseal",
}