# 部署文档

### 部署后端
创建文件kubez.yaml,然后kubectl apply -f 它
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubez-configmap
  namespace: default
data:
  config.yaml: |-
    logs:
      level: debug
      timeformat: 2006/1/02 15:04:05
      showcaller: false
    program:
      # 后端程序运行的端口号
      port: 8080
      # 
      jwtsecret: jwtHello
      username: admin
      password: 123
      kubeconfig: 你的base64格式kubeconfig
       
immutable: false
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubez
  labels:
    app: kubez
  annotations:
    app: kubez
  namespace: default
spec:
  selector:
    matchLabels:
      app: kubez
  replicas: 1
  template:
    metadata:
      labels:
        app: kubez
      annotations:
        app: kubez
    spec:
      restartPolicy: Always
      dnsPolicy: ClusterFirst
      volumes:
        - name: kubez-config
          configMap:
            name: kubez-configmap
            defaultMode: 420
            optional: true
      containers:
        - name: kubez
          image: reactive0/kubez:lastest
          imagePullPolicy: IfNotPresent
          volumeMounts:
            - name: kubez-config
              mountPath: /kubez/config.yaml
              subPath: config.yaml

---
apiVersion: v1
kind: Service
metadata:
  name: kubez
  namespace: default
spec:
  ports:
    - port: 8080
      targetPort: 8080
      protocol: TCP
      name: kubez
  selector:
    app: kubez
  type: ClusterIP
  sessionAffinity: None
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    app: kubez
  name: admin
  namespace: kubez
type: Opaque
stringData:
  annotations: bWFwWzEyMzo0NTZd
  clusterconfig: ====
  clusterid: cluster1
  clustername: 集群1
  version: v1.28
```


### 部署前端
同上
```
apiVersion: v1
kind: Service
metadata:
  name: kubez-frontend
  namespace: default
  labels:
    app: kubez-frontend
spec:
  ports:
    - name: kubez-frontend
      protocol: TCP
      port: 8080
      targetPort: 80
      nodePort: 30261
  selector:
    app: kubez-frontend
  type: NodePort
  sessionAffinity: None

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubez-frontend
  namespace: default
  labels:
    app: kubez-frontend
  annotations:
    app: kubez-frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubez-frontend
  template:
      labels:
        app: kubez-frontend
      annotations:
        app: kubez-frontend
    spec:
      containers:
        - name: kubez-frontend
          image: reactive0/kubez_frontend:lastest
          ports:
            - name: kubez-frontend
              containerPort: 80
              protocol: TCP
          imagePullPolicy: Always
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
```

之后访问nodeport对应的端口,再在集群管理中添加集群即可,kubeconfig同样是base64格式

添加完集群后如果不能访问,请重启后端服务的pod
